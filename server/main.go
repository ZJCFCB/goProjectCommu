package main

import (
	"server/model/dao"
	"server/view"
	"time"
)

func main() {
	dao.InitPool("localhost:6379", 8, 0, 300*time.Second) //初始化redis连接池
	dao.InitUserDao()
	var s view.EnterServer
	_ = s.Run()
}
