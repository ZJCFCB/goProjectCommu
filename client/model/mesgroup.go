package model

type MesGroup struct {
	Toall string `json:"toall"`
	Id    int    `json:"id`
	Name  string `json:"name"`
}

type MesGroupRes struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}
type MesGroupInform struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
	Toall   string `json:"toall"`
	Idfrom  int    `json:"idfrom"`
	Name    string `json:"name"`
}
