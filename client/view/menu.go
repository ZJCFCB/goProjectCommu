package view

import (
	"client/controller"
	"client/util"
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
			var password string
			fmt.Printf("请输入用户名：")
			fmt.Scanln(&id)
			fmt.Printf("请输入密码：")
			fmt.Scanln(&password)

			//实例化一个用户管理类，可以通过它进行对用户的相关操作
			var up *controller.UserProcess = &controller.UserProcess{}

			//请求服务器建立连接（8889端口） 用于登录校验等

			//Todo 全局变量？连接池？

			up.MakeConn("localhost:8889")
			defer up.Conn.Close() //记得退出的时候关闭连接

			up.Tf = &util.Transfer{Conn: up.Conn}

			//根据用户输入的账号密码进行登录校验
			ok, err, name := up.LoginCheck(id, password)
			up.Id = id
			up.Name = name
			if ok {
				fmt.Println("登录成功")
				//登录成果后用户在这里面完成登陆后的操作 比如聊天
				var us *controller.Userserve = &controller.Userserve{
					Conn: up.Conn,
					Id:   id,
					Tf:   &util.Transfer{Conn: up.Conn},
					Name: name,
				}
				us.ServerProcessMessage()
			} else {
				fmt.Println("用户登录失败", err)
			}

		case 2:
			//注册之后是需要重新登录的
			var id int
			var passwd, name string

			fmt.Printf("请输入用户名：")
			fmt.Scanln(&id)
			fmt.Printf("请输入密码：")
			fmt.Scanln(&passwd)
			fmt.Printf("请输入姓名：")
			fmt.Scanln(&name)

			var up *controller.UserProcess = &controller.UserProcess{}

			up.MakeConn("localhost:8889") //与服务器建立连接（8889端口）
			defer up.Conn.Close()         //记得退出的时候关闭连接
			up.Tf = &util.Transfer{Conn: up.Conn}
			ok, err := up.Regist(id, passwd, name)

			if ok {
				fmt.Println("用户注册成功，可以退出登录")
			} else {
				fmt.Println("用户注册失败", err)
			}
		case 3:
			loop = util.Exit()
		default:
			fmt.Println("输入错误")
		}
		if loop {
			break
		}
	}
}
