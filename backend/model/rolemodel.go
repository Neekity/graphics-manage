package model

import (
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	RoleModel interface {
		FindOne(id int) (*Role, error)
		List(name string) ([]Role, error)
		UpdateOrCreate(id int, materialInfo Role) (*Role, error)
		Delete(id uint) error
	}

	defaultRoleModel struct {
		GormDB *gorm.DB
		table  string
	}

	Role struct {
		ID        uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement;comment:主键id;" json:"id"`
		Name      string `gorm:"type:string;comment:角色名称;size:64;not null;" json:"name"`
		Key       string `gorm:"type:string;uniqueIndex;comment:角色key;size:64;not null;" json:"key"`
		CreatedAt helper.MyTime
		UpdatedAt helper.MyTime
	}
)

func (m *defaultRoleModel) Delete(id uint) error {
	return m.GormDB.Delete(&Role{ID: id}).Error
}

func (m *defaultRoleModel) UpdateOrCreate(id int, RoleInfo Role) (*Role, error) {
	var role Role
	if err := m.GormDB.Model(&Role{}).Where("id = ?", id).Assign(RoleInfo).FirstOrCreate(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (m *defaultRoleModel) List(name string) ([]Role, error) {
	var roles []Role
	query := m.GormDB.Table(m.table)
	if name != "" {
		query.Scopes(helper.QueryKey("%"+name+"%", "name", "like"))
	}
	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (m *defaultRoleModel) FindOne(id int) (*Role, error) {
	var role Role
	if err := m.GormDB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func NewRoleModel(gdb *gorm.DB) RoleModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '角色表'  charset = utf8mb4;").
		AutoMigrate(&Role{})
	return &defaultRoleModel{
		GormDB: gdb,
		table:  "`Role`",
	}
}
