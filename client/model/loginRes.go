package model

// 登录返回结果
type LoginRes struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}
