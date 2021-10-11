package config

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/rest"
	"golang.org/x/oauth2"
)

type Config struct {
	rest.RestConf
	DataSource    string
	FrontUrl      string
	BeanstalkConf []dq.Beanstalk
	oauth2.Config
	UserInfoUrl string
	JwtAuth     struct {
		AccessSecret string
		AccessExpire int64
	}
}
