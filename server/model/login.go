package model

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginRes struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
	Name    string `json:"name`
}
