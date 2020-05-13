<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [ccyb](#ccyb)
    - [现在开始](#现在开始)
    - [文件分层](#文件分层)
    - [输出格式统一封装](#输出格式统一封装)
    - [定义中间件链路日志打印](#定义中间件链路日志打印)
    - [请求数据绑定到结构体与校验](#请求数据绑定到结构体与校验)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# ccyb
Gin best practices, gin development scaffolding, too late to explain, get on the bus.

使用gin构建了企业级脚手架，代码简洁易读，可快速进行高效web开发。
主要功能有：
1. 请求链路日志打印，涵盖mysql/redis/request_in/request_out
2. 支持多语言错误信息提示及自定义错误提示。
3. 支持了多配置环境
4. 封装了 log日志 /redis /mysql / http.client 常用方法

项目地址：
### 现在开始
- 安装软件依赖
go mod使用请查阅：

https://blog.csdn.net/e421083458/article/details/89762113
```
git clone 
cd ccyb
go mod tidy
```
- 运行脚本
```
go run main.go

➜  ccyb git:(master) ✗ go run main.go
------------------------------------------------------------------------
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /register                 --> test/controller/auth.Register (6 handlers)
[GIN-debug] POST   /login                    --> test/controller/auth.Login (6 handlers)
[GIN-debug] POST   /user-info                --> test/controller/auth.UserInfo (7 handlers)
2020/05/13 09:43:50  [INFO] HttpServerRun::8880
```
创建用户表：
```
curl -X POST 'http://127.0.0.1:8880/login' -d 'username=cc&password=cccc'
{
    "code": 0,
    "msg": "",
    "data": {
        "email": "cc@qq.com",
        "id": "1",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImNjIiwiZW1haWwiOiJjY0BxcS5jb20iLCJleHAiOjE1ODk0MjA2NjgsImlhdCI6MTU4OTMzNDI2OH0.N2BVVHrsvKbUjaabIv6VeJml7aUXhxEhS4nSnmMU-vA",
        "username": "cc"
    },
    "trace_id": "7f0000015ebb50fcb27845f0658221b0",
    "stack": null
}

查看链路日志（确认是不是一次请求查询，都带有相同trace_id）：
tail -f gin_scaffold.inf.log
[INFO][2020-05-13T09:54:04.275][log.go:43] _com_request_in||method=POST||body=----------------------------393153290403551399013061\r\nContent-Disposition: form-data; name=\"username\"\r\n\r\ncc\r\n----------------------------393153290403551399013061\r\nContent-Disposition: form-data; name=\"password\"\r\n\r\ncccc\r\n----------------------------393153290403551399013061--\r\n||spanid=21bb533d380704bb||traceid=7f0000015ebb533cfc2c4f70104dc7b0||cspanid=||uri=/login||args=map[]||from=127.0.0.1
[INFO][2020-05-13T09:54:04.275][log.go:43] _com_request_out||from=127.0.0.1||traceid=7f0000015ebb533cfc2c4f70104dc7b0||spanid=21bb533d380704bb||uri=/login||method=POST||args=map[password:[cccc] username:[cc]]||response={\"code\":0,\"msg\":\"\",\"data\":{\"email\":\"cc@qq.com\",\"id\":\"1\",\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImNjIiwiZW1haWwiOiJjY0BxcS5jb20iLCJleHAiOjE1ODk0MjEyNDQsImlhdCI6MTU4OTMzNDg0NH0.jppAjkuYXTsL1IXkQduY7MgKEIljR_9ZbjT-hSmhnDY\",\"username\":\"cc\"},\"trace_id\":\"7f0000015ebb533cfc2c4f70104dc7b0\",\"stack\":null}||proc_time=0.0049855||cspanid=

```
- 测试参数绑定与多语言验证

```
curl -X POST 'http://127.0.0.1:8880/login' -d 'username=cc'
{
    "code": 101,
    "msg": "密码为必填字段",
    "data": "",
    "trace_id": "7f0000015ebb542ea08803b8658221b0",
    "stack": "test/common.DefaultGetValidParams\n\tC:/cf/test/ginweb/common/params.go:37\ntest/request/auth.(*LoginInput).ParseLoginParams\n\tC:/cf/test/ginweb/request/auth/authRequestParamsBuilder.go:45\ntest/controller/auth.Login\n\tC:/cf/test/ginweb/controller/auth/login.go:16\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/context.go:156\ntest/middleware.RecoveryMiddleware.func1\n\tC:/cf/test/ginweb/middleware/recovery.go:33\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/context.go:156\ntest/middleware.TranslationMiddleware.func1\n\tC:/cf/test/ginweb/middleware/translation.go:73\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/context.go:156\ntest/middleware.RequestLog.func1\n\tC:/cf/test/ginweb/middleware/request_log.go:65\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/context.go:156\ngithub.com/gin-gonic/gin.RecoveryWithWriter.func1\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/recovery.go:83\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/context.go:156\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/logger.go:241\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/context.go:156\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/gin.go:409\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n\tC:/cf/gocode/pkg/mod/github.com/gin-gonic/gin@v1.6.2/gin.go:367\nnet/http.serverHandler.ServeHTTP\n\tC:/cf/Go/src/net/http/server.go:2774\nnet/http.(*conn).serve\n\tC:/cf/Go/src/net/http/server.go:1878\nruntime.goexit\n\tC:/cf/Go/src/runtime/asm_amd64.s:1337"
}

curl -X POST 'http://127.0.0.1:8880/login' -d 'username=cc&lang=en'
{
    "code": 101,
    "msg": "Password is a required field",
    "data": "",
    "trace_id": "7f0000015ebb693c625c479c5a8581b0",
    "stack": ""
}
```

### 文件分层
```
├── README.md
├── common 公共文件夹
│   ├── log.go     日志记录
│   └── params.go  参数校验
├── conf           配置文件夹
│   ├── api.go
│   └── config.go   初始化配置
├── controller      控制器
│   └── auth
├── errorCode 自定义错误码
│   └── errorCode.go
├── model   数据模型层
├── helper  逻辑处理层
├── lib     相当于自定义package
├── logs    日志
├── request 请求参数处理层
├── router  路由层
├── static  静态文件层
├── middleware  中间件
│   ├── ip_auth.go
│   ├── jwt_auth.go
│   ├── recovery.go
│   ├── request_log.go
│   ├── response.go
│   └── translation.go
├── go.mod go module管理文件
│   └──go.sum
├── example.yaml    配置文件示例
└── main.go
```

### 输出格式统一封装
```
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceId: traceId}
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
```
### 定义中间件链路日志打印
```
package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)
//链路请求日志
func RequestInLog(c *gin.Context) {
	traceContext := lib.NewTrace()
	if traceId := c.Request.Header.Get("com-header-rid"); traceId != "" {
		traceContext.TraceId = traceId
	}
	if spanId := c.Request.Header.Get("com-header-spanid"); spanId != "" {
		traceContext.SpanId = spanId
	}
	c.Set("startExecTime", time.Now())
	c.Set("trace", traceContext)
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back

	lib.Log.TagInfo(traceContext, "_com_request_in", map[string]interface{}{
		"uri":    c.Request.RequestURI,
		"method": c.Request.Method,
		"args":   c.Request.PostForm,
		"body":   string(bodyBytes),
		"from":   c.ClientIP(),
	})
}
//链路输出日志
func RequestOutLog(c *gin.Context) {
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")
	startExecTime, _ := st.(time.Time)
	public.ComLogNotice(c, "_com_request_out", map[string]interface{}{
		"uri":       c.Request.RequestURI,
		"method":    c.Request.Method,
		"args":      c.Request.PostForm,
		"from":      c.ClientIP(),
		"response":  response,
		"proc_time": endExecTime.Sub(startExecTime).Seconds(),
	})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		defer RequestOutLog(c)
		isMatched := false
		for _, host := range lib.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched{
			ResponseError(c, InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
```
### 请求数据绑定到结构体与校验

dto/demo.go
```
package request

import (
	"github.com/gin-gonic/gin"
	"test/common"
)

type LoginInput struct {
	UserName 	string `json:"username" form:"username" comment:"姓名" validate:"required"`
	Password	string `json:"password" form:"password" comment:"密码" validate:"required"`
}


func (params *LoginInput) ParseLoginParams(c *gin.Context) error{
	//检查时候格式是否正确
	if err := common.DefaultGetValidParams(c, params); err != nil{
		return  err
	}

	return nil
}
```
controller/demo.go
```
        //参数校验
	loginInputParams := &request.LoginInput{}
	err := loginInputParams.ParseLoginParams(c)

	if err != nil {
		middleware.ResponseError(c, errorCode.ErrParamFormat, err)
		return
	}
```
<!-- ### log日志 常用方法
参考文档：https://github.com/e421083458/golang_common -->
