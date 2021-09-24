package model

import (
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	GraphicsCasbinRuleModel interface {
		DeleteRole(role string, roleType string) error
		GetOwnerChannels(userId int) ([]int, error)
	}

	defaultGraphicsCasbinRuleModel struct {
		GormDB *gorm.DB
		table  string
	}

	GraphicsCasbinRule struct {
		ID        uint          `json:"id" gorm:"primaryKey;column:id;type:int(10) unsigned auto_increment"`
		Ptype     string        `json:"ptype" gorm:"column:ptype;type:varchar(255)"`
		V0        string        `json:"v0" gorm:"column:v0;type:varchar(255)"`
		V1        string        `json:"v1" gorm:"column:v1;type:varchar(255)"`
		V2        string        `json:"v2" gorm:"column:v2;type:varchar(255)"`
		V3        string        `json:"v3" gorm:"column:v3;type:varchar(255)"`
		V4        string        `json:"v4" gorm:"column:v4;type:varchar(255)"`
		V5        string        `json:"v5" gorm:"column:v5;type:varchar(255)"`
		CreatedAt helper.MyTime `json:"created_at" gorm:"column:created_at;type:timestamp"`
		UpdatedAt helper.MyTime `json:"updated_at" gorm:"column:updated_at;type:timestamp"`
	}
)

func (m *defaultGraphicsCasbinRuleModel) GetOwnerChannels(userId int) ([]int, error) {
	return []int{1, 3, 4, 5}, nil
}

func (m *defaultGraphicsCasbinRuleModel) DeleteRole(role string, roleType string) error {
	panic("implement me")
}

func NewGraphicsCasbinRuleModel(gdb *gorm.DB) GraphicsCasbinRuleModel {
	return &defaultGraphicsCasbinRuleModel{
		GormDB: gdb,
		table:  "`graphics_casbin_rule`",
	}
}
