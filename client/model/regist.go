package model

type RegistMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegistRes struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}
