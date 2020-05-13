package controller

import (
	"cy/errorCode"
	"cy/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserInfo(c *gin.Context) {

	userId, _ := strconv.Atoi(c.PostForm("user_id"))

	claims := c.MustGet("claims").(*middleware.CustomClaims)

	if userId != claims.ID {
		middleware.ResponseError(c, errorCode.ErrParamFormat, errors.New("user_id 参数错误"))
		return
	}

	middleware.ResponseSuccess(c, claims)
	return

	/*userId, _:= strconv.Atoi(c.PostForm("user_id"))

	infos, err := helper.Userinfo(userId)

	if err != nil {
		middleware.ResponseError(c, errorCode.ErrParamFormat, err)
		return
	}

	middleware.ResponseSuccess(c, infos)
	return*/
}
