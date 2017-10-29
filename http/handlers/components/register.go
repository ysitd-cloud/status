package components

import "github.com/gin-gonic/gin"

var serviceMap = make(map[string]string)

func init() {
	serviceMap["account"] = "account-manager"
	serviceMap["status"] = "status-api"
}

func Register(group *gin.RouterGroup) {
	for endpoint, service := range serviceMap {
		group.GET("/"+endpoint, createComponentStatusEndpoint(service))
	}
}
