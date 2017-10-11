package cdn

import "github.com/gin-gonic/gin"

func Register(app *gin.RouterGroup) {
	app.GET("/health", handle)
}
