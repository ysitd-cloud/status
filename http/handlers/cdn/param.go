package cdn

import (
	"os"
)

var cfEmail string = os.Getenv("CF_EMAIL")

var cfKey string = os.Getenv("CF_KEY")

var cfZoneID string = os.Getenv("CF_ZONE_ID")
