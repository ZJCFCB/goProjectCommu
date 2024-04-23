package controller

import (
	"fmt"
	"net"
	"server/model"
	"server/util"
)

type BaseProcess struct {
	Conn net.Conn
}

func (B *BaseProcess) ServerProcessMes(mes *model.Message) (err error) { // 根据消息类型的不同，调用不同的处理函数
	switch mes.Type {
	case util.LoginMesType:
		// 处理登录的相关信息
		up := &UserProcess{Conn: B.Conn}
		up.HandLogin(mes)
		err = nil
	default:
	}
	return
}

func (B *BaseProcess) Process() {
	for {
		tf := &util.Transfer{Conn: B.Conn}
		mess, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ReadPkg failed")
			return
		}
		fmt.Println("message is ", mess)
		B.ServerProcessMes(&mess)
	}
}
