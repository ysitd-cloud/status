package components

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/k8s-utils/services"
)

func createComponentStatusEndpoint(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := services.GetServiceStatus(coreV1, namespace, name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, &resp)
			c.Abort()
		}
	}
}
