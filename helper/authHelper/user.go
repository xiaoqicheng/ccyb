package authHelper

import (
	"cy/config"
	"cy/model/authModel"
	"errors"
	"time"
)

/**
@desc 用户注册
*/
func Save(user authModel.User) int {

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	config.LocalMysql.Table("t_bi_user").Create(&user)

	return user.Id
}

func Userinfo(userId int) (userInfo []authModel.User, err error) {

	userInfo = authModel.UserInfo(userId)

	if len(userInfo) == 0 {
		return userInfo, errors.New("id参数错误")
	}

	return userInfo, nil
}
