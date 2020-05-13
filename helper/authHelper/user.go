package helper

import (
	"cy/config"
	model "cy/model/auth"
	"errors"
	"time"
)

/**
@desc 用户注册
*/
func Save(user model.User) int {

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	config.LocalMysql.Table("t_bi_user").Create(&user)

	return user.Id
}

func Userinfo(userId int) (userInfo []model.User, err error) {

	userInfo = model.UserInfo(userId)

	if len(userInfo) == 0 {
		return userInfo, errors.New("id参数错误")
	}

	return userInfo, nil
}
