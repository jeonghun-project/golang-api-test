package sample

import (
	"test/controller/response"

	"github.com/valyala/fasthttp"
)

type checkRes struct {
	Success bool `json:"success"`
}

func Get(ctx *fasthttp.RequestCtx) {

	res := &checkRes{Success: true}

	response.JSON(ctx, res)
	return
}
