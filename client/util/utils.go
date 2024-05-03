package util

import "fmt"

func Exit() bool {
	var confir string
	fmt.Printf("你确定要退出系统吗？请输入Y或y确认：")
	fmt.Scanln(&confir)
	if confir == "Y" || confir == "y" {
		return true
	} else {
		return false
	}
}
