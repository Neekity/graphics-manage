package model

import (
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	GraphicsMaterialModel interface {
		FindOne(id int) (*GraphicsMaterial, error)
		List(name string, Type string) ([]GraphicsMaterial, error)
		UpdateOrCreate(id int, materialInfo GraphicsMaterial) (*GraphicsMaterial, error)
	}

	defaultGraphicsMaterialModel struct {
		GormDB *gorm.DB
		table  string
	}

	GraphicsMaterial struct {
		ID        uint          `json:"id" gorm:"primaryKey;column:id;type:int(10) unsigned auto_increment"`
		URL       string        `json:"url" gorm:"column:url;type:varchar(256);not null;comment:素材图片地址"`                                    // 素材图片地址
		Name      string        `json:"name" gorm:"column:name;type:varchar(128);not null;comment:素材名称"`                                    // 素材名称
		Type      string        `json:"type" gorm:"column:type;type:varchar(16);not null;comment:素材类型"`                                     // 素材类型
		Content   string        `json:"content" gorm:"column:content;type:longtext;not null;comment:素材内容"`                                  // 素材内容
		CreatedAt helper.MyTime `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"` // 创建时间
		UpdatedAt helper.MyTime `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
		Abstract  string        `json:"abstract" gorm:"column:abstract;type:varchar(256);comment:摘要"`                                       // 摘要
		ChannelID int           `json:"channel_id" gorm:"column:channel_id;type:int(10) unsigned;not null;comment:频道id"`                    // 频道id
	}
)

func (m *defaultGraphicsMaterialModel) UpdateOrCreate(id int, materialInfo GraphicsMaterial) (*GraphicsMaterial, error) {
	var material GraphicsMaterial
	if err := m.GormDB.Model(&GraphicsMaterial{}).Where("id = ?", id).Assign(materialInfo).FirstOrCreate(&material).Error; err != nil {
		return nil, err
	}
	return &material, nil
}

func (m *defaultGraphicsMaterialModel) List(name string, Type string) ([]GraphicsMaterial, error) {
	var materials []GraphicsMaterial
	query := m.GormDB.Table(m.table)
	query.Scopes(helper.QueryKey(Type, "type", "="))
	if name != "" {
		query.Scopes(helper.QueryKey("%"+name+"%", "name", "like"))
	}
	if err := query.Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}

func (m *defaultGraphicsMaterialModel) FindOne(id int) (*GraphicsMaterial, error) {
	var material GraphicsMaterial
	if err := m.GormDB.First(&material, id).Error; err != nil {
		return nil, err
	}
	return &material, nil
}

func NewGraphicsMaterialModel(gdb *gorm.DB) GraphicsMaterialModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '图文素材表'  charset = utf8mb4;").
		AutoMigrate(&GraphicsMaterial{})
	return &defaultGraphicsMaterialModel{
		GormDB: gdb,
		table:  "`graphics_material`",
	}
}
