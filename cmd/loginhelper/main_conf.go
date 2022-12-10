package main

import (
	"encoding/json"
	"path/filepath"
	"pbbot_app_loginhelper/pkg/dto"
	"sync"
	"time"

	"gitee.com/lyhuilin/log"
	"gitee.com/lyhuilin/util"
)

var mux sync.RWMutex

func getPathConf() (retText string) {
	path := filepath.Join(util.Getwd())
	if !util.IsExist(path) {
		util.MkDir(path)
	}
	path = filepath.Join(path, "conf", "loginhelper.json")
	return path
}
func Init_Conf() {
	path := getPathConf()
	if !util.IsExist(path) {
		LoginJsonCreate()
	}
	bytes, err := util.ReadFile(path)
	if err != nil {
		log.Errorf(err, "读取配置文件 %s 出错了", path)
	}
	err = json.Unmarshal(bytes, Conf)
	if err != nil {
		log.Errorf(err, "加载配置文件 %s 出错了", path)
	}
}

func isZero(botId int64) (retBool bool) {
	switch {
	case botId == 0:
		retBool = true
	case botId == int64(1234567890):
		retBool = true
	case botId <= 10000:
		retBool = true
	default:
	}
	return
}
func LoginJsonCreate() {
	login := &dto.CreateBotReq{
		BotId:          1234567890,
		Password:       "123456ab",
		DeviceSeed:     1234567890,
		ClientProtocol: 1,
	}

	var logins LoginHelperConf
	logins.Logins = append(logins.Logins, login)
	logins.Logins = append(logins.Logins, login)
	logins.ServerURL = "http://127.0.0.1:9000"
	logins.CheckSleep = 30 * time.Minute

	outBody, err := json.MarshalIndent(&logins, "", "\t")
	if err != nil {
		log.Errorf(err, "序列化配置信息出错(json.MarshalIndent): %v", logins)
		return
	}
	path := getPathConf()
	err = util.WriteFile(path, outBody)
	if err != nil {
		log.Errorf(err, "写入配置信息到文件(%s)出错(util.WriteFile)", path)
	} else {
		log.Infof("写入配置信息到文件(%s):%s", path, string(outBody))
	}
}
