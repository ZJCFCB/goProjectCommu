package view

import (
	"fmt"
)

type Conutrol struct{}

func (c *Conutrol) Run() {

	var key int
	var loop bool = false
	for {
		fmt.Println("---------------------欢迎登录多人聊天系统---------------------")
		fmt.Println("\t\t\t1.登录聊天室")
		fmt.Println("\t\t\t2.注册用户")
		fmt.Println("\t\t\t3.退出系统")
		fmt.Printf("请选择(1-3) : ")
		fmt.Scanln(&key)

		switch key {
		case 1:
			var name int
			var password string
			fmt.Printf("请输入用户名：")
			fmt.Scanln(&name)
			fmt.Printf("请输入密码：")
			fmt.Scanln(&password)
			if LoginCheck(name, password) == nil {
				fmt.Println("登录成功")
			} else {
				fmt.Println("账户名或密码错误")
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
