package main

import (
	"encoding/json"
	"pbbot_app_loginhelper/pkg/dto"

	"path/filepath"
	"sync"
	"time"

	"gitee.com/lyhuilin/log"
	"gitee.com/lyhuilin/util"
)

var mux sync.RWMutex

// var feedGroupList map[int64]bool
// var toGroupList map[int64]bool
// var botIdList map[int64]bool

func getPathConf() (retText string) {
	path := filepath.Join(util.Getwd())
	if !util.IsExist(path) {
		util.MkDir(path)
	}
	path = filepath.Join(path, "conf", "loginhelper.json")
	return path
}
func Init_Conf() {
	// mux.Lock()
	// defer mux.Unlock()
	// feedGroupList = make(map[int64]bool, 0)
	// toGroupList = make(map[int64]bool, 0)
	// botIdList = make(map[int64]bool, 0)
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
	logins.CheckSleep = 10 * time.Minute

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