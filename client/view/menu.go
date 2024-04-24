package view

import (
	"client/controller"
	"fmt"
)

type EnterClient struct{}

func (c *EnterClient) Run() {

	var key int
	var loop bool = false
	for {
		//这里算是一个控制界面了
		fmt.Println("---------------------欢迎登录多人聊天系统---------------------")
		fmt.Println("\t\t\t1.登录聊天室")
		fmt.Println("\t\t\t2.注册用户")
		fmt.Println("\t\t\t3.退出系统")
		fmt.Printf("请选择(1-3) : ")
		fmt.Scanln(&key)

		switch key {
		case 1:
			//处理用户登录的相关信息
			var id int
			var password, name string
			fmt.Printf("请输入用户名：")
			fmt.Scanln(&id)
			fmt.Printf("请输入密码：")
			fmt.Scanln(&password)
			fmt.Printf("请输入昵称：")
			fmt.Scanln(&name)
			var up *controller.UserProcess = &controller.UserProcess{}

			up.MakeConn("localhost:8889") //与服务器建立连接（8889端口）
			defer up.Conn.Close()         //记得退出的时候关闭连接  （其实我觉得这个连接应该作为一个全局变量）

			ok, err := up.LoginCheck(id, password, name) //根据用户输入的账号密码进行登录校验
			if ok {
				fmt.Println("登录成功")

				controller.ShowLoginMenu()

				go controller.ServerProcessMessage(up.Conn) //开启一个协程保持通讯
			} else {
				fmt.Println("用户登录失败", err)
			}

		case 2:
			fmt.Println("注册成功")
		case 3:
			fmt.Println("退出系统")
			loop = true
		default:
			fmt.Println("输入错误")
		}
		if loop {
			break
		}
	}
}
