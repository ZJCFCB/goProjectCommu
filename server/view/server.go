package view

import (
	"fmt"
	"net"
	"server/controller"
)

type Server struct {
}

func (s *Server) Run() {
	s.Register()
}

func (s *Server) Register() {
	//先监听8889端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil { //监听失败
		fmt.Println("监听失败")
		return
	}

	for {
		fmt.Println("waitting...")
		conn, err := listen.Accept()
		if err != nil { //accepte失败
			fmt.Println("accepte失败")
			return
		}
		//连接成功，启动协程与客户端保持通讯
		go Ganhuo(conn)
	}
}

func Ganhuo(conn net.Conn) {
	defer conn.Close()
	pro := controller.BaseProcess{Conn: conn}
	pro.Process()
}
