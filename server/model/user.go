package model

type User struct {
	UserId   int64  `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
