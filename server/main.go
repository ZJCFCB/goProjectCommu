package main

import (
	"fmt"
	"server/view"
)

func main() {
	fmt.Println("开始干活了。。。")
	var s view.Server
	s.Run()
}
