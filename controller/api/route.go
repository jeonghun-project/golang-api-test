package api

import (
	"test/config"
	"test/controller/middleware/logMiddleware"
	"test/controller/response"
	"test/controller/sample"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func InitRouter() fasthttp.RequestHandler {
	return middleware(routes().Handler)
}

func middleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	if !config.IsProductionMode() {
		h = logMiddleware.Handler(h)
	}
	h = fasthttp.CompressHandlerBrotliLevel(h, fasthttp.CompressBrotliBestCompression, fasthttp.CompressBestCompression)
	return h
}

func routes() *fasthttprouter.Router {
	r := fasthttprouter.New()

	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		response.Text(ctx, "success")
	})

	r.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		response.Text(ctx, "pong")
	})

	addRoute(r)
	return r
}

func addRoute(r *fasthttprouter.Router) {

	r.GET("/sample", sample.Get)
}
