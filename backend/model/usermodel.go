package model

import (
	"github.com/casbin/casbin/v2"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
	"strconv"
)

type (
	UserModel interface {
		FindOne(id int) (*User, error)
		List(name string) ([]User, error)
		UpdateOrCreate(id int, materialInfo User) (*User, error)
		Delete(id uint) error
		GetUserRoles(id uint, enforce *casbin.Enforcer) ([]Role, error)
		Assign(id uint, casbinRoles []string, enforce *casbin.Enforcer) error
	}

	defaultUserModel struct {
		GormDB *gorm.DB
		table  string
	}

	User struct {
		ID        uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement;comment:主键id;" json:"id"`
		Name      string `gorm:"type:string;comment:姓名;size:64;not null;" json:"name"`
		Email     string `gorm:"type:string;uniqueIndex;comment:邮箱;size:64;not null;" json:"email"`
		Mobile    string `gorm:"type:string;comment:手机号;size:16;not null;" json:"mobile"`
		Avatar    string `gorm:"type:string;comment:头像;size:256;not null;" json:"avatar"`
		CreatedAt helper.MyTime
		UpdatedAt helper.MyTime
		DeletedAt gorm.DeletedAt
	}
)

func (m *defaultUserModel) Assign(id uint, casbinRoles []string, enforce *casbin.Enforcer) error {
	err := m.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := enforce.LoadPolicy(); err != nil {
			return err
		}
		if _, err := enforce.RemoveFilteredNamedGroupingPolicy(constant.CasbinDefaultRole, 0, strconv.Itoa(int(id))); err != nil {
			return err
		}
		var rules [][]string
		for _, role := range casbinRoles {
			rules = append(rules, []string{strconv.Itoa(int(id)), role})
		}

		if _, err := enforce.AddNamedGroupingPolicies(constant.CasbinDefaultRole, rules); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (m *defaultUserModel) GetUserRoles(id uint, enforce *casbin.Enforcer) ([]Role, error) {
	var roles []Role
	if err := enforce.LoadPolicy(); err != nil {
		return nil, err
	}
	policies := enforce.GetFilteredNamedGroupingPolicy(constant.CasbinDefaultRole, 0, strconv.Itoa(int(id)))
	for _, policy := range policies {
		var role Role

		if err := m.GormDB.Model(&Role{}).Where("casbin_role = ?", policy[1]).First(&role).Error; err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}
	return roles, nil
}

func (m *defaultUserModel) Delete(id uint) error {
	return m.GormDB.Delete(&User{ID: id}).Error
}

func (m *defaultUserModel) UpdateOrCreate(id int, userInfo User) (*User, error) {
	var material User
	if err := m.GormDB.Model(&User{}).Where("id = ?", id).Assign(userInfo).FirstOrCreate(&material).Error; err != nil {
		return nil, err
	}
	return &material, nil
}

func (m *defaultUserModel) List(name string) ([]User, error) {
	var users []User
	query := m.GormDB.Table(m.table)
	if name != "" {
		query.Scopes(helper.QueryKey("%"+name+"%", "name", "like"))
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *defaultUserModel) FindOne(id int) (*User, error) {
	var material User
	if err := m.GormDB.First(&material, id).Error; err != nil {
		return nil, err
	}
	return &material, nil
}

func NewUserModel(gdb *gorm.DB) UserModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '用户表'  charset = utf8mb4;").
		AutoMigrate(&User{})
	return &defaultUserModel{
		GormDB: gdb,
		table:  "`user`",
	}
}
