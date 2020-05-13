/**
@date 2020-04-28
@desc 返回中间件
*/
package middleware

import (
	"cy/config"
	"cy/lib"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type ResponseCode int

//返回码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode
	InvalidRequestErrorCode ResponseCode = 401
)

type Response struct {
	ErrorCode ResponseCode `json:"code"`
	ErrorMsg  string       `json:"msg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"trace_id"`
	Stack     interface{}  `json:"stack"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	stack := ""
	if c.Query("is_debug") == "1" || config.ApiConfig.HostEnv == "debug" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceId: traceId, Stack: stack}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
