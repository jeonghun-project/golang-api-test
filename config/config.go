package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Port = ""

var ServerType = ""

const Dev = "DEVELOPMENT"

const AppName = "golang-fasthttp-test"

func init() {
	godotenv.Load()

	Port = os.Getenv("PORT")

	ServerType = os.Getenv("TYPE")

	if Port == "" {
		if ServerType == Dev {
			Port = "4000"
		} else {
			Port = "80"
		}
	}
}

func IsProductionMode() bool {
	return strings.EqualFold(ServerType, "PRODUCTION")
}
