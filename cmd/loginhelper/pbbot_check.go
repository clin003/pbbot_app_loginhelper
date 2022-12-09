package main

import (
	"time"

	pbbot "github.com/2mf8/GoPbBot"
)

func taskCheckPbbotOnline() {
	time.Sleep(60 * time.Second)
	for {
		getOnlineBotList()
		time.Sleep(Conf.CheckSleep)
	}
}

func getOnlineBotList() {
	mux.Lock()
	defer mux.Unlock()
	list := make(map[int64]bool, 0)

	for botId, _ := range pbbot.Bots {
		list[botId] = true
	}

	for _, i := range Conf.Logins {
		if _, ok := list[i.BotId]; !ok {
			if isZero(i.BotId) {
				continue
			}
			postPbbotCreate(Conf.ServerURL, i)
			time.Sleep(30 * time.Second)
		}
	}
}
