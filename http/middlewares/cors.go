package middlewares

import "github.com/gin-gonic/gin"

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
