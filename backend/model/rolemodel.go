package model

import (
	"github.com/casbin/casbin/v2"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
	"strconv"
)

type (
	RoleModel interface {
		FindOne(id uint, enforce *casbin.Enforcer) (*RoleInfo, error)
		List(name string) ([]Role, error)
		UpdateOrCreate(id uint, RoleInfo Role, userIds []uint, enforce *casbin.Enforcer) (*Role, error)
		Delete(id uint) error
	}

	defaultRoleModel struct {
		GormDB *gorm.DB
		table  string
	}

	Role struct {
		ID         uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement;comment:主键id;" json:"id"`
		Name       string `gorm:"type:string;comment:角色名称;size:64;not null;" json:"name"`
		CasbinRole string `gorm:"type:string;uniqueIndex;comment:角色key;size:64;not null;" json:"casbin_role"`
		CreatedAt  helper.MyTime
		UpdatedAt  helper.MyTime
	}

	RoleInfo struct {
		Role
		Users []User `json:"users"`
	}
)

func (m *defaultRoleModel) Delete(id uint) error {
	return m.GormDB.Delete(&Role{ID: id}).Error
}

func (m *defaultRoleModel) UpdateOrCreate(id uint, RoleInfo Role, userIds []uint, enforce *casbin.Enforcer) (*Role, error) {
	var role Role

	err := m.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := m.GormDB.Model(&Role{}).Where("id = ?", id).Assign(RoleInfo).FirstOrCreate(&role).Error; err != nil {
			return err
		}
		if err := enforce.LoadPolicy(); err != nil {
			return err
		}
		if _, err := enforce.RemoveFilteredNamedGroupingPolicy(constant.CasbinDefaultRole, 1, role.CasbinRole); err != nil {
			return err
		}
		var rules [][]string
		for _, userId := range userIds {
			rules = append(rules, []string{strconv.Itoa(int(userId)), role.CasbinRole})
		}

		if _, err := enforce.AddNamedGroupingPolicies(constant.CasbinDefaultRole, rules); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
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

func (m *defaultRoleModel) FindOne(id uint, enforce *casbin.Enforcer) (*RoleInfo, error) {
	var role Role
	var users []User
	if err := m.GormDB.First(&role, id).Error; err != nil {
		return nil, err
	}
	if err := enforce.LoadPolicy(); err != nil {
		return nil, err
	}
	policies := enforce.GetFilteredNamedGroupingPolicy(constant.CasbinDefaultRole, 1, role.CasbinRole)
	for _, policy := range policies {
		var user User
		userId, err := strconv.Atoi(policy[0])
		if err != nil {
			return nil, err
		}
		if err := m.GormDB.First(&user, userId).Error; err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	roleInfo := RoleInfo{
		role,
		users,
	}

	return &roleInfo, nil
}

func NewRoleModel(gdb *gorm.DB) RoleModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '角色表'  charset = utf8mb4;").
		AutoMigrate(&Role{})
	return &defaultRoleModel{
		GormDB: gdb,
		table:  "`role`",
	}
}
