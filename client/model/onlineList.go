package model

type OnlineListMes struct {
	UserId int `json:"userId"`
}

type OnlineListRes struct {
	Errno      int    `json:"errno"`
	Message    string `json:"message"`
	OnlineList []int  `json:"onlineList"`
}
