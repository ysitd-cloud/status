package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/status/http/handlers/cdn"
	"github.com/ysitd-cloud/status/http/handlers/components"
	"github.com/ysitd-cloud/status/http/handlers/traffic"
	"github.com/ysitd-cloud/status/http/middlewares"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.CORS)
	app.Use(middlewares.Security)

	{
		group := app.Group("/traffic")
		traffic.Register(group)
	}

	{
		group := app.Group("/cdn")
		cdn.Register(group)
	}

	{
		group := app.Group("/components")
		components.Register(group)
	}
}
