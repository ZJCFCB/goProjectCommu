package model

import (
	"errors"
)

var (
	ERROR_USER_NOTEXIT = errors.New("用户不存在")
	ERROR_user_IS_EXIT = errors.New("用户已经存在，请直接登录")
	ERROR_PASSWD_RONG  = errors.New("密码错误")
)
