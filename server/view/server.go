package view

import (
	"net"
	"server/controller"
	"server/util"
)

type EnterServer struct {
}

func (s *EnterServer) Run() (err error) {
	//服务器监听8889端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		return util.ERROR_LISTERN_FAILED
	}

	for {
		var conn net.Conn
		conn, err = listen.Accept()
		if err != nil { //accepte失败
			return util.ERROR_ACCEPT_FAILED
		}
		//连接成功，启动协程与客户端保持通讯
		go Communication(conn)
	}
}

// 这里会开启一个协程，用于保持跟用户的通讯
// 用BaseProcess来控制与当前连接用户的通讯
func Communication(conn net.Conn) {
	defer conn.Close()
	pro := controller.BaseProcess{Conn: conn}
	_ = pro.Process()
}
