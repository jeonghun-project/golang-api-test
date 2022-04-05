package response

import (
	"encoding/json"
	"fmt"
	"strings"
	"test/config"
	"test/util/customLog"
	"test/util/errorFactory"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	StatusSuccess   = 200
	ContentTypeJSON = "application/json"
)

type Response struct {
	Status int         `json:"status"` //200 means success, anything else error
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func newResponse(status int, data interface{}, err string) *Response {
	return &Response{Status: status, Data: data, Error: err}
}

func Text(ctx *fasthttp.RequestCtx, data string) {
	ctx.SetContentType("text/plain")
	ctx.Response.Header.Set("server", config.AppName)
	_, _ = fmt.Fprintf(ctx, "%s", data)
}

func HTML(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetContentType("text/html")
	ctx.Response.Header.Set("server", config.AppName)
	_, _ = fmt.Fprintln(ctx, data)
}

func JSON(ctx *fasthttp.RequestCtx, data interface{}) {
	sendJsonResponse(ctx, newResponse(StatusSuccess, data, ""))
}

func Success(ctx *fasthttp.RequestCtx) {
	sendJsonResponse(ctx, newResponse(StatusSuccess, map[string]interface{}{"success": true}, ""))
}

func ServerError(ctx *fasthttp.RequestCtx, err error) {
	sendJsonResponse(ctx, newResponse(errorFactory.UnknownServerError, nil, ""))
	log.WithFields(customLog.Fields("response", "ServerError", string(ctx.Request.URI().Path()))).Error(err)
}

func OtherError(ctx *fasthttp.RequestCtx, errorCode int, msg string) {
	sendJsonResponse(ctx, newResponse(errorCode, nil, msg))
	log.WithFields(customLog.Fields("response", "OtherError", string(ctx.Request.URI().Path()))).Error(errorCode, msg)
}

func SomethingMissingError(ctx *fasthttp.RequestCtx, msg string) {
	OtherError(ctx, errorFactory.SomethingMissing, msg)
}

func UnauthorizedError(ctx *fasthttp.RequestCtx) {
	sendJsonResponse(ctx, newResponse(errorFactory.Unauthorized, nil, "unauthorized"))
	log.WithFields(customLog.Fields("response", "Unauthorized", string(ctx.Request.URI().Path()))).Error("")
}

func sendJsonResponse(ctx *fasthttp.RequestCtx, r *Response) {
	ctx.SetContentType(ContentTypeJSON)
	ctx.Response.Header.Set("server", config.AppName)
	if strings.Contains(string(ctx.RequestURI()), "/v3/") {
		_ = json.NewEncoder(ctx).Encode(r)
	} else {
		//To keep supporting older versions.
		status := r.Status
		data := r.Data
		errMessage := r.Error

		if status == StatusSuccess {
			if data == nil {
				data = map[string]interface{}{"success": true}
			}
			_ = json.NewEncoder(ctx).Encode(data)
			return
		}
		ctx.Response.SetStatusCode(500)
		oldRes := map[string]interface{}{
			"message": errMessage,
			"code":    status,
			"success": false,
		}
		_ = json.NewEncoder(ctx).Encode(oldRes)
		return
	}
}
