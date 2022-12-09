package main

import (
	"pbbot_app_loginhelper/pkg/dto"
	"time"
)

// 登录助手配置模型
type LoginHelperConf struct {
	Debug      bool          `yaml:"debug"  json:"debug"`
	ServerURL  string        `yaml:"gmc_server_url" json:"gmc_server_url"`
	CheckSleep time.Duration `yaml: "check_sleep" json:"check_sleep"`

	Logins []*dto.CreateBotReq
}

var Conf *LoginHelperConf

func init() {
	Conf = &LoginHelperConf{}
}
