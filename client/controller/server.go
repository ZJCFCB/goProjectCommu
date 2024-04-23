package controller

import (
	"client/util"
	"fmt"
	"net"
)

//显示登录成功的界面

func ShowLoginMenu() {
	var key int
	fmt.Println("---------------------登录成功---------------------")
	fmt.Println("\t\t\t1.显示在线用户列表")
	fmt.Println("\t\t\t2.发送消息")
	fmt.Println("\t\t\t3.信息列表")
	fmt.Println("\t\t\t4.退出登录")
	fmt.Printf("请选择(1-4) : ")
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出登录")
	default:
		fmt.Println("重新输入")
	}
}

func ServerProcessMessage(conn net.Conn) {
	//创建一个transfer实例，不停地读取服务器信息
	tf := &util.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端等待读取服务器发送过来的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		fmt.Println("mes = ", mes)
	}
}
