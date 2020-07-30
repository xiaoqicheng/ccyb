package authRequest

import (
	"cy/common"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" validate:"required"`
	Email    string `json:"email" form:"email" comment:"邮箱" validate:"required"`
}

func (params *RegisterInput) ParseRegisterParams(c *gin.Context) error {

	//检查时候格式是否正确
	if err := common.DefaultGetValidParams(c, params); err != nil {
		return err
	}

	return nil
}

/**
@desc 登录接口 参数格式校验
*/
type LoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" validate:"required"`
}

func (params *LoginInput) ParseLoginParams(c *gin.Context) error {
	//检查时候格式是否正确
	if err := common.DefaultGetValidParams(c, params); err != nil {
		return err
	}

	return nil
}
