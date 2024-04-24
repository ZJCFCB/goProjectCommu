package util

import "errors"

const (
	Success = 200 + iota
	NoRegistered
	PasswdIsWrong
)

const (
	SERVICE_HAS_WRONG = 500
)

var (
	ERROR_USER_NOTEXIT      = errors.New("用户不存在")
	ERROR_user_IS_EXIT      = errors.New("用户已经存在，请直接登录")
	ERROR_PASSWD_RONG       = errors.New("密码错误")
	ERROR_LISTERN_FAILED    = errors.New("监听失败")
	ERROR_ACCEPT_FAILED     = errors.New("建立连接失败/accept失败")
	ERROR_MARSHAL_FAILED    = errors.New("marshal failed")
	ERROR_UN_MARSHAL_FAILED = errors.New("unmarshal failed")
	ERROR_WRITE_CONN_FAILED = errors.New("write to conn failed")
	ERROR_READ_CONN_FAILED  = errors.New("read from conn failed")
)
