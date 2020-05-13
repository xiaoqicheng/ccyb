/**
@date 2020-04-28
@desc IP白名单中间件
*/

package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range viper.GetStringMap("http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			ResponseError(c, InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
