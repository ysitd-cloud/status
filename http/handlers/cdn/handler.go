package cdn

import (
	"net/http"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

var api *cloudflare.API

func init() {
	client, err := cloudflare.New(cfKey, cfEmail)
	if err != nil {
		panic(err)
	}
	api = client
}

func handle(c *gin.Context) {
	defer c.Abort()
	if data, err := api.ZoneAnalyticsDashboard(cfZoneID, cloudflare.ZoneAnalyticsOptions{}); err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
	} else {
		c.JSON(http.StatusOK, data.Totals)
	}
}
