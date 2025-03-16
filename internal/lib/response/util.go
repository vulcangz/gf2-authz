package response

import (
	"encoding/json"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func ReturnError(r *ghttp.Request, statusCode int, err error) error {
	RJson(r, statusCode, err.Error())
	r.Exit()

	return nil
}

func ReturnError1(r *ghttp.Request, statusCode int, err error) error {
	r.Response.WriteStatus(statusCode)

	r.Response.WriteJson(Response{
		Code:      statusCode,
		Error:     true,
		Message:   err.Error(),
		Timestamp: gtime.Timestamp(),
		TraceID:   gctx.CtxId(r.Context()),
	})

	return nil
}

func ReturnHTTPError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)

	responseBytes, err := json.Marshal(Response{
		Code:      statusCode,
		Error:     true,
		Message:   err.Error(),
		Timestamp: gtime.Timestamp(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(responseBytes)
}
