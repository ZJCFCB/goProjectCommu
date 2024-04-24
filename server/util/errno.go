package util

import "errors"

const (
	Success = 200 + iota
	NoRegistered
	PasswdIsWrong
	UserHasExist
)

const (
	SERVICE_HAS_WRONG = 500
)

var (
	ERROR_USER_NOTEXIT           = errors.New("用户不存在")
	ERROR_user_IS_EXIST          = errors.New("用户已经存在，请直接登录")
	ERROR_PASSWD_RONG            = errors.New("密码错误")
	ERROR_USER_HAS_EXIT          = errors.New("注册失败，该用户已经存在")
	ERROR_LISTERN_FAILED         = errors.New("监听失败")
	ERROR_ACCEPT_FAILED          = errors.New("建立连接失败/accept失败")
	ERROR_MARSHAL_FAILED         = errors.New("marshal failed")
	ERROR_UN_MARSHAL_FAILED      = errors.New("unmarshal failed")
	ERROR_WRITE_CONN_FAILED      = errors.New("write to conn failed")
	ERROR_READ_CONN_FAILED       = errors.New("read from conn failed")
	ERROR_PUSH_TO_REDIS_FAILED   = errors.New("写入redis失败")
	ERROR_READ_FROM_REDIS_FAILED = errors.New("从redis中读取失败")
)
