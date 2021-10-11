package model

import (
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	MenuModel interface {
		MenuList(permissionNames []string) ([]MenuTree, error)
		UpdateOrCreateUser(id uint, menuInfo Menu) (*Menu, error)
	}

	defaultMenuModel struct {
		table  string
		GormDB *gorm.DB
	}

	Menu struct {
		ID        uint           `gorm:"column:id;primaryKey" json:"id,omitempty"`
		Name      string         `gorm:"column:name" json:"name,omitempty"`           // 菜单名
		ParentId  int64          `gorm:"column:parent_id" json:"parent_id,omitempty"` // 父级菜单
		Icon      string         `gorm:"column:icon" json:"icon,omitempty"`           // 图标
		Path      string         `gorm:"column:path" json:"path,omitempty"`           // 前端路由
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
		CreatedAt helper.MyTime  `json:"created_at"`
		UpdatedAt helper.MyTime  `json:"updated_at"`
	}

	MenuTree struct {
		Children []Menu `json:"children"`
		Menu
	}
)

func (m *defaultMenuModel) UpdateOrCreateUser(id uint, menuInfo Menu) (*Menu, error) {
	var menu Menu
	m.GormDB.Model(&Menu{}).Where("id = ?", id).Assign(menuInfo).FirstOrCreate(&menu)
	return &menu, nil
}

func (m *defaultMenuModel) MenuList(permissionNames []string) ([]MenuTree, error) {
	var parentMenus []Menu
	var childrenMenus []Menu
	m.GormDB.Model(&Menu{}).Where("parent_id = ?", 0).Find(&parentMenus)
	m.GormDB.Model(&Menu{}).Where("`Path` IN ?", permissionNames).Find(&childrenMenus)
	var menuTrees []MenuTree
	for _, parentMenu := range parentMenus {
		var curChildrenMenus []Menu
		for _, childrenMenu := range childrenMenus {
			if int64(parentMenu.ID) == childrenMenu.ParentId {
				curChildrenMenus = append(curChildrenMenus, childrenMenu)
			}
		}
		if helper.InArray(parentMenu.Path, permissionNames) || len(curChildrenMenus) > 0 {
			menuTrees = append(menuTrees, MenuTree{curChildrenMenus, parentMenu})
		}
	}
	return menuTrees, nil
}

func NewMenuModel(gdb *gorm.DB) MenuModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '菜单表'  charset = utf8mb4;").
		AutoMigrate(&Menu{})
	return &defaultMenuModel{
		table:  "`menu`",
		GormDB: gdb,
	}
}
