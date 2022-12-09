package main

import (
	"gitee.com/lyhuilin/log"

	pbbot "github.com/2mf8/GoPbBot"
	// "github.com/2mf8/GoPbBot/proto_gen/onebot"
)

func Init_Handle() {
	pbbot.HandleConnect = func(bot *pbbot.Bot) {
		log.Infof("[连接] 新机器人已连接：%d", bot.BotId)
		if !Conf.Debug {
			return
		}
		log.Infof("[已连接] 所有机器人列表：")
		for botId, _ := range pbbot.Bots {
			log.Infof("[已连接]: %d", botId)
		}

		if groupList, err := bot.GetGroupList(); err != nil {
			log.Errorf(err, "b.GetGroupList()")
		} else {
			for _, group := range groupList.Group {
				log.Infof("[INFO] Bot(%v) GroupID: %d Name: %s", bot.BotId, group.GroupId, group.GroupName)
			}
		}

	}
}
