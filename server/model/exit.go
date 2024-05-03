package model

type ExitMes struct {
	UserId int `json:"userId"`
}

type ExitResMes struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}
