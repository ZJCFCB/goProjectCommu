package model

type MesSide struct {
	Side   string `json:"side"`
	ToId   int    `json:"toid`
	MyId   int    `json:"myid`
	MyName string `json:"myname"`
}

type MesSideRes struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}

type MesSideInform struct {
	Errno    int    `json:"errno"`
	Message  string `json:"message"`
	Side     string `json:"side"`
	Idfrom   int    `json:"idfrom"`
	Namefrom string `json:"namefrom"`
}
