package traffic

import (
	"fmt"
	"os"
)

var traefikAddress string = os.Getenv("TRAEFIK_ENDPOINT")

var healthEndpoint string = fmt.Sprintf("http://%s/health", traefikAddress)
