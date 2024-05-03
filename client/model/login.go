package model

type LoginMes struct {
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
}

// 登录返回结果
type LoginRes struct {
	Errno         int    `json:"errno"`
	Message       string `json:"message"`
	OnlineUserIds []int  `json:"onlineUserIds"`
}
