package main

import (
	"fmt"
	"server/model/dao"
	"server/view"
	"time"
)

func main() {
	fmt.Println("开始干活了。。。")
	dao.InitPool("localhost:6379", 8, 0, 300*time.Second) //初始化redis连接池
	dao.InitUserDao()
	var s view.Server
	s.Run()
}
