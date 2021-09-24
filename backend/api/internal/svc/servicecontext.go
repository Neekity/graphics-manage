package svc

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/tal-tech/go-zero/rest"

	"go-project/graphics-manage/backend/api/internal/config"
	"go-project/graphics-manage/backend/api/internal/middleware"
	"go-project/graphics-manage/backend/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config                       config.Config
	AccessLog                    rest.Middleware
	CheckDataPermission          rest.Middleware
	Auth                         rest.Middleware
	GraphicsCasbinRuleModel      model.GraphicsCasbinRuleModel
	GraphicsCasbinRuleEnforce    *casbin.Enforcer
	GraphicsChannelModel         model.GraphicsChannelModel
	GraphicsMaterialModel        model.GraphicsMaterialModel
	GraphicsMessageReceiverModel model.GraphicsMessageReceiverModel
	GraphicsMessageModel         model.GraphicsMessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	gdb, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	a, err := gormadapter.NewAdapterByDBUseTableName(gdb, "graphics", "casbin_rule")
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("/data/resource/rbac_with_resource_roles_model.conf", a)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                       c,
		AccessLog:                    middleware.NewAccessLogMiddleware().Handle,
		CheckDataPermission:          middleware.NewCheckDataPermissionMiddleware().Handle,
		Auth:                         middleware.NewAuthMiddleware().Handle,
		GraphicsCasbinRuleEnforce:    e,
		GraphicsChannelModel:         model.NewGraphicsChannelModel(gdb),
		GraphicsMaterialModel:        model.NewGraphicsMaterialModel(gdb),
		GraphicsMessageReceiverModel: model.NewGraphicsMessageReceiverModel(gdb),
		GraphicsMessageModel:         model.NewGraphicsMessageModel(gdb),
	}
}
