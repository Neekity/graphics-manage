package model

import (
	"fmt"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	GraphicsChannelModel interface {
		ChannelOwner(channelIds []int) ([]GraphicsChannel, error)
		ChannelList(name string) ([]GraphicsChannel, error)
		UpdateOrCreate(id int, channelInfo GraphicsChannel) (*GraphicsChannel, error)
		FindOne(id int) (*GraphicsChannel, error)
	}

	defaultGraphicsChannelModel struct {
		GormDB *gorm.DB
		table  string
	}

	GraphicsChannel struct {
		ID             uint64        `json:"id" gorm:"primaryKey;column:id;type:bigint(20) unsigned auto_increment"`
		Name           string        `json:"name" gorm:"column:name;type:varchar(256);not null;comment:频道名称"`                                    // 频道名称
		CreatedAt      helper.MyTime `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"` // 创建时间
		UpdatedAt      helper.MyTime `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
		AudienceConfig string        `json:"audience_config" gorm:"column:audience_config;type:longtext;not null;comment:频道听众配置"`                // 频道听众配置
	}
)

func (m *defaultGraphicsChannelModel) FindOne(id int) (*GraphicsChannel, error) {
	var channel GraphicsChannel
	if err := m.GormDB.First(&channel, id).Error; err != nil {
		return nil, err
	}
	return &channel, nil
}

func (m *defaultGraphicsChannelModel) UpdateOrCreate(id int, channelInfo GraphicsChannel) (*GraphicsChannel, error) {
	var channel GraphicsChannel
	err := m.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&GraphicsChannel{}).Where("id = ?", id).Assign(channelInfo).FirstOrCreate(&channel).Error; err != nil {
			return err
		}
		if id == 0 {
			roles := []Role{
				{
					Name:       fmt.Sprintf(constant.ChannelWriteRoleName, channel.Name),
					CasbinRole: fmt.Sprintf(constant.CasbinChannelWriteRole, channel.ID),
				},
				{
					Name:       fmt.Sprintf(constant.ChannelReadRoleName, channel.Name),
					CasbinRole: fmt.Sprintf(constant.CasbinChannelReadRole, channel.ID),
				},
			}
			if err := tx.Create(&roles).Error; err != nil {
				return err
			}
			casbinModels := []GraphicsCasbinRule{
				{
					Ptype: constant.CasbinPolicy,
					V0:    fmt.Sprintf(constant.CasbinChannelWriteRole, channel.ID),
					V1:    fmt.Sprintf(constant.CasbinChannelMaterialResourceRole, channel.ID),
					V2:    constant.CasbinPermissionWrite,
				},
				{
					Ptype: constant.CasbinPolicy,
					V0:    fmt.Sprintf(constant.CasbinChannelReadRole, channel.ID),
					V1:    fmt.Sprintf(constant.CasbinChannelMaterialResourceRole, channel.ID),
					V2:    constant.CasbinPermissionRead,
				},
				{
					Ptype: constant.CasbinPolicy,
					V0:    fmt.Sprintf(constant.CasbinChannelWriteRole, channel.ID),
					V1:    fmt.Sprintf(constant.CasbinChannelMessageResourceRole, channel.ID),
					V2:    constant.CasbinPermissionWrite,
				},
				{
					Ptype: constant.CasbinPolicy,
					V0:    fmt.Sprintf(constant.CasbinChannelReadRole, channel.ID),
					V1:    fmt.Sprintf(constant.CasbinChannelMessageResourceRole, channel.ID),
					V2:    constant.CasbinPermissionRead,
				},
			}
			if err := tx.Create(&casbinModels).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (m *defaultGraphicsChannelModel) ChannelList(name string) ([]GraphicsChannel, error) {
	var channels []GraphicsChannel
	query := m.GormDB.Table(m.table)

	if name != "" {
		query.Scopes(helper.QueryKey("%"+name+"%", "name", "like"))
	}
	if err := query.Find(&channels).Error; err != nil {
		return nil, err
	}
	return channels, nil
}

func (m *defaultGraphicsChannelModel) ChannelOwner(channelIds []int) ([]GraphicsChannel, error) {
	var channels []GraphicsChannel

	if err := m.GormDB.Table(m.table).Find(&channels, channelIds).Error; err != nil {
		return nil, err
	}
	return channels, nil
}

func NewGraphicsChannelModel(gdb *gorm.DB) GraphicsChannelModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '图文频道表'  charset = utf8mb4;").
		AutoMigrate(&GraphicsChannel{})
	return &defaultGraphicsChannelModel{
		GormDB: gdb,
		table:  "`graphics_channel`",
	}
}
