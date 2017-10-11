package cdn

import (
	"fmt"
	"os"
)

var cfEmail string = os.Getenv("CF_EMAIL")

var cfKey string = os.Getenv("CF_KEY")

var cfZoneID string = os.Getenv("CF_ZONE_ID")

var endpointUrl string = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/analytics/dashboard", cfZoneID)
