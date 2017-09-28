package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/status/http"
)

func main() {
	app := gin.Default()

	http.Register(app)

	app.Run(os.Getenv("LISTEN_ADDRESS"))
}
