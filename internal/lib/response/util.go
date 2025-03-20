package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func ReturnError(r *ghttp.Request, statusCode int, err error, message string) error {
	r.Response.WriteStatus(statusCode)

	text := ""
	if message == "" || message == err.Error() {
		text = err.Error()
	} else {
		text = fmt.Sprintf("%s: %v", message, err)
	}

	r.Response.WriteJson(Response{
		Code:      statusCode,
		Error:     true,
		Message:   text,
		Timestamp: gtime.Timestamp(),
		TraceID:   gctx.CtxId(r.Context()),
	})

	return nil
}

func RExitError(r *ghttp.Request, statusCode int, err error) error {
	RJson(r, statusCode, err.Error())
	r.Exit()

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
