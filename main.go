package main

import (
	"fmt"
	"test/config"
	"test/controller/api"

	"github.com/lab259/cors"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func init() {}

func main() {
	var router fasthttp.RequestHandler

	router = api.InitRouter()

	log.Infoln("Listening on port", config.Port)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%s", config.Port), cors.AllowAll().Handler(router)))
}
