package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/status/http/handlers/traffic"
	"github.com/ysitd-cloud/status/http/middlewares"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.CORS)

	{
		group := app.Group("/traffic")
		traffic.Register(group)
	}
}