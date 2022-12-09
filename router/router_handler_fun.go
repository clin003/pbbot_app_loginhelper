package router

import (
	. "pbbot_app_loginhelper/handler"
	"pbbot_app_loginhelper/pkg/constvar"

	"gitee.com/lyhuilin/pkg/errno"
	"github.com/gin-gonic/gin"
)

// 404 Not found
func ApiNotFound(c *gin.Context) {
	SendResponse(c, errno.Err404, constvar.APPDesc404())
}

// API Hello
func ApiHello(c *gin.Context) {
	SendResponse(c, errno.SayHello, constvar.APPDesc())
}

// API ping
func ApiPing(c *gin.Context) {
	SendResponse(c, errno.PONG, constvar.APP_VERSION)
}
