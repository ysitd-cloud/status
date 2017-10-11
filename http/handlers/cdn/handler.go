package cdn

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func handle(c *gin.Context) {
	defer c.Abort()

	request := gorequest.New()
	var data cfResponse

	_, _, erros := request.Get(endpointUrl).
		Set("Content-Type", "application/json").
		Set("X-Auth-Email", cfEmail).
		Set("X-Auth-Key", cfKey).
		EndStruct(&data)

	if len(erros) > 0 {
		c.AbortWithError(http.StatusBadGateway, erros[0])
		return
	}

	c.JSON(http.StatusOK, data.Result.Totals)

}
