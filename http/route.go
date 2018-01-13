package http

import (
	"github.com/gin-gonic/gin"
	"code.ysitd.cloud/component/status/http/handlers/cdn"
	"code.ysitd.cloud/component/status/http/handlers/components"
	"code.ysitd.cloud/component/status/http/handlers/traffic"
	"code.ysitd.cloud/component/status/http/middlewares"
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
