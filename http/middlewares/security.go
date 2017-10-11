package middlewares

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

var Security gin.HandlerFunc = secure.New(secure.Config{
	FrameDeny:          true,
	BrowserXssFilter:   true,
	ContentTypeNosniff: true,
	STSSeconds:         5184000,
})
