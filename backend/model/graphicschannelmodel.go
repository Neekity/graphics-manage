package model

import (
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type (
	GraphicsChannelModel interface {
		GetChannels(channelIds []int) ([]GraphicsChannel, error)
		ChannelList(name string) ([]GraphicsChannel, error)
		UpdateOrCreate(id int, channelInfo GraphicsChannel, owners []ChannelOwner, enforce *casbin.Enforcer) (*GraphicsChannel, error)
		FindOne(id int, enforce *casbin.Enforcer) (*ChannelDetail, error)
		GetOwnerChannels(userId string, enforcer *casbin.Enforcer) ([]int, error)
		GetUserChannels(userId string) ([]int, error)
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

	ChannelDetail struct {
		ChannelOwners string `json:"channel_owners"`
		Name          string `json:"name"`
		Audiences     string `json:"audiences"`
	}

	ChannelOwner struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (m *defaultGraphicsChannelModel) GetUserChannels(userId string) ([]int, error) {
	var results []string
	if err := m.GormDB.Model(&GraphicsCasbinRule{}).Where("ptype = ?", constant.CasbinPolicy).
		Where("v0 = ?", userId).Pluck("v1", &results).Error; err != nil {
		return nil, err
	}
	var channelIds []int
	for _, item := range results {
		subject := strings.Split(item, constant.CasbinPermissionSymbol)
		if len(subject) != 3 || !strings.Contains(item, constant.CasbinMessagePolicy[:len(constant.CasbinChannelRole)-2]) {
			continue
		}
		messageId, err := strconv.Atoi(subject[2])
		if err != nil {
			return nil, err
		}
		var message GraphicsMessage
		if err := m.GormDB.Model(&GraphicsMessage{}).First(&message, messageId).Error; err != nil {
			continue
		}
		channelIds = append(channelIds, message.ChannelID)
	}
	return channelIds, nil
}

func (m *defaultGraphicsChannelModel) GetOwnerChannels(userId string, enforcer *casbin.Enforcer) ([]int, error) {
	var results []int
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	policy := enforcer.GetFilteredNamedGroupingPolicy(constant.CasbinDefaultRole,
		0, userId)
	for _, item := range policy {
		role := strings.Split(item[1], constant.CasbinPermissionSymbol)
		if len(role) != 3 || !strings.Contains(item[1], constant.CasbinChannelRole[:len(constant.CasbinChannelRole)-2]) {
			continue
		}
		channelId, err := strconv.Atoi(role[2])
		if err != nil {
			return nil, err
		}
		results = append(results, channelId)
	}
	return results, nil
}

func (m *defaultGraphicsChannelModel) FindOne(id int, enforce *casbin.Enforcer) (*ChannelDetail, error) {
	var channel GraphicsChannel
	if err := m.GormDB.First(&channel, id).Error; err != nil {
		return nil, err
	}
	if err := enforce.LoadPolicy(); err != nil {
		return nil, err
	}
	policy := enforce.GetFilteredNamedGroupingPolicy(constant.CasbinDefaultRole,
		1, fmt.Sprintf(constant.CasbinChannelRole, channel.ID))
	var channelOwners []ChannelOwner
	for _, item := range policy {
		var user User
		userId, err := strconv.Atoi(item[0])
		if err != nil {
			return nil, err
		}
		m.GormDB.Model(&User{}).First(&user, userId)
		channelOwners = append(channelOwners, ChannelOwner{
			Id:   userId,
			Name: user.Name,
		})
	}
	channelOwnerByte, err := json.Marshal(channelOwners)
	if err != nil {
		return nil, err
	}
	channelDetail := ChannelDetail{
		Audiences:     channel.AudienceConfig,
		Name:          channel.Name,
		ChannelOwners: string(channelOwnerByte),
	}
	return &channelDetail, nil
}

func (m *defaultGraphicsChannelModel) UpdateOrCreate(id int, channelInfo GraphicsChannel, owners []ChannelOwner, enforce *casbin.Enforcer) (*GraphicsChannel, error) {
	var channel GraphicsChannel
	err := m.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&GraphicsChannel{}).Where("id = ?", id).Assign(channelInfo).FirstOrCreate(&channel).Error; err != nil {
			return err
		}
		if id == 0 {
			role := Role{
				Name:       fmt.Sprintf(constant.ChannelRoleName, channel.Name),
				CasbinRole: fmt.Sprintf(constant.CasbinChannelRole, channel.ID),
			}
			if err := tx.Create(&role).Error; err != nil {
				return err
			}
			casbinModels := []GraphicsCasbinRule{
				{
					Ptype: constant.CasbinPolicy,
					V0:    fmt.Sprintf(constant.CasbinChannelRole, channel.ID),
					V1:    fmt.Sprintf(constant.CasbinChannelMaterialResourceRole, channel.ID),
					V2:    constant.CasbinPermissionWrite,
				},
				{
					Ptype: constant.CasbinPolicy,
					V0:    fmt.Sprintf(constant.CasbinChannelRole, channel.ID),
					V1:    fmt.Sprintf(constant.CasbinChannelMessageResourceRole, channel.ID),
					V2:    constant.CasbinPermissionWrite,
				},
			}
			if err := tx.Create(&casbinModels).Error; err != nil {
				return err
			}
		}
		if err := enforce.LoadPolicy(); err != nil {
			return err
		}
		if _, err := enforce.RemoveFilteredNamedGroupingPolicy(constant.CasbinDefaultRole,
			1, fmt.Sprintf(constant.CasbinChannelRole, channel.ID)); err != nil {
			return err
		}
		for _, item := range owners {
			if _, err := enforce.AddNamedGroupingPolicy(constant.CasbinDefaultRole, []string{strconv.Itoa(item.Id), fmt.Sprintf(constant.CasbinChannelRole, channel.ID)}); err != nil {
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

func (m *defaultGraphicsChannelModel) GetChannels(channelIds []int) ([]GraphicsChannel, error) {
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
