package pbbot_wsserver

import (
	"gitee.com/lyhuilin/log"
	pbbot "github.com/2mf8/GoPbBot"
	"github.com/gin-gonic/gin"
)

func PbBotWs(c *gin.Context) {
	if err := pbbot.UpgradeWebsocket(c.Writer, c.Request); err != nil {
		log.Errorf(err, "[失败] 创建机器人失败")
	}
}
