package authModel

import (
	"cy/config"
)

type User struct {
	Id        int    `json:"id" gorm:"column:id"`
	UserName  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	Email     string `json:"email" gorm:"column:email"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at"" description:"更新时间"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at" description:"创建时间"`
}

func (f *User) TableName() string {
	return "t_bi_user"
}

/**
@desc 根据用户名检查用户是否存在
*/

func IExist(username string) bool {

	var user []User

	config.LocalMysql.Table("t_bi_user").Where("username = ?", username).Find(&user)

	if len(user) == 0 {
		return false
	}

	return true
}

/**
@desc 根据用户名和密码查询是否存在
*/
func IsRight(username string, password string) (user []User, ok bool) {

	config.LocalMysql.Table("t_bi_user").Select("id, username, email").Where("username = ? and password = ?", username, password).Find(&user)

	if len(user) == 0 {
		return user, false
	}

	return user, true
}

func UserInfo(userId int) (user []User) {

	config.LocalMysql.Table("t_bi_user").Select("id, username, email").Where("id = ?", userId).Find(&user)

	return user
}
