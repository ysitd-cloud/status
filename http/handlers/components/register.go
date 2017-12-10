package components

import "github.com/gin-gonic/gin"

var serviceMap = map[string]string{
	"account": "account-manager",
	"status": "status-api",
	"controller": "app-controller",
}

func Register(group *gin.RouterGroup) {
	for endpoint, service := range serviceMap {
		group.GET("/"+endpoint, createComponentStatusEndpoint(service))
	}
}
