package logMiddleware

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/valyala/fasthttp"
)

func Handler(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		next(ctx)
		end := time.Now()

		ruri := string(ctx.RequestURI())
		if ruri == "/ping" || ruri == "/v1/lot_sources" {
			return
		}
		log.Infoln(string(ctx.Method()), "\t", string(ctx.RequestURI()), "\t", end.Sub(start), "\t", ctx.Response.StatusCode())
	}
}
