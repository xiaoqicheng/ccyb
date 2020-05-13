package controller

import (
	"cy/errorCode"
	authHelper "cy/helper/authHelper"
	"cy/middleware"
	model "cy/model/auth"
	request "cy/request/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	registerInputParams := &request.RegisterInput{}
	var user model.User
	if err := registerInputParams.ParseRegisterParams(c); err != nil {
		middleware.ResponseError(c, errorCode.ErrParamFormat, err)
		return
	} else {
		user.UserName = registerInputParams.UserName
		user.Password = registerInputParams.Password
		user.Email = registerInputParams.Email
	}

	// 校验用户名是否存在
	if ok := model.IExist(registerInputParams.UserName); ok {
		middleware.ResponseError(c, errorCode.ErrParamFormat, errors.New("用户名已存在，请登录"))
		return
	}

	userId := authHelper.Save(user)

	//注册成功返回用户信息 跳转至登录界面
	output := make(map[string]string)
	output["id"] = strconv.Itoa(userId)
	output["name"] = user.UserName

	middleware.ResponseSuccess(c, output)
	return
}
