package traffic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func getHealth(c *gin.Context) {
	req := gorequest.New()
	_, body, errs := req.Get(healthEndpoint).End()

	if len(errs) > 0 {
		c.AbortWithError(http.StatusInternalServerError, errs[0])
	} else {
		c.Data(http.StatusOK, "application/json", []byte(body))
		c.Abort()
	}
}
