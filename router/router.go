package router

import (
	"pbbot_app_loginhelper/handler/pbbot_wsserver"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)

	g.NoRoute(ApiNotFound)
	g.GET("/", ApiHello)
	g.GET("/ping", ApiPing)

	// WebSocket 路由
	wserver := g.Group("/ws")
	{
		wserver.GET("/rq/", pbbot_wsserver.PbBotWs)
		wserver.GET("/cq/", pbbot_wsserver.PbBotWs)
	}

	return g
}
