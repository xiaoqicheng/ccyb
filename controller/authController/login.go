package authController

import (
	"cy/errorCode"
	"cy/middleware"
	"cy/model/authModel"
	"cy/request/authRequest"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Login(c *gin.Context) {
	//参数校验
	loginInputParams := &authRequest.LoginInput{}
	err := loginInputParams.ParseLoginParams(c)

	if err != nil {
		middleware.ResponseError(c, errorCode.ErrParamFormat, err)
		return
	}

	// 校验用户名是否存在
	if ok := authModel.IExist(loginInputParams.UserName); !ok {
		middleware.ResponseError(c, errorCode.ErrParamFormat, errors.New("无效用户名"))
		return
	}

	//校验密码是否正确
	userInfo, ok := authModel.IsRight(loginInputParams.UserName, loginInputParams.Password)
	if !ok {
		middleware.ResponseError(c, errorCode.ErrParamFormat, errors.New("密码不正确"))
		return
	}

	//生成token
	var customClaime middleware.CustomClaims
	for _, values := range userInfo {
		customClaime.ID = values.Id
		customClaime.Name = values.UserName
		customClaime.Email = values.Email
	}

	jwt := middleware.NewJWT()

	token, err := jwt.CreateToken(customClaime)

	if err != nil {
		middleware.ResponseError(c, errorCode.ErrTokenCreate, err)
		return
	}

	//返回用户登录信息
	var output = make(map[string]string)
	for _, items := range userInfo {
		output["id"] = strconv.Itoa(items.Id)
		output["username"] = items.UserName
		output["email"] = items.Email
		output["token"] = token
	}

	middleware.ResponseSuccess(c, output)
	return
}
