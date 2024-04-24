package controller

import (
	"net"
	"server/model"
	"server/util"
)

/*
流程控制类
在与客户端建立连接后，这里主要负责保持通讯并处理客户请求
*/
type BaseProcess struct {
	Conn net.Conn
}

func (B *BaseProcess) ServerProcessMes(mes *model.Message) (err error) { // 根据消息类型的不同，调用不同的处理函数
	switch mes.Type {
	case util.LoginMesType:
		// 处理登录的相关信息
		up := &UserProcess{Conn: B.Conn}
		err = up.HandLogin(mes)
	case util.RegistMesType:
		up := &UserProcess{Conn: B.Conn}
		err = up.HandRegist(mes)
	default:
	}
	return err
}

// 处理信息的入口
func (B *BaseProcess) Process() (err error) {
	for {
		tf := &util.Transfer{Conn: B.Conn}
		mess, err := tf.ReadPkg()
		if err != nil {
			return err
		}
		//收到了来自客户端的消息，交给ServerProcessMes处理
		//这里面的错误处理要根据类型判断是否为服务器内部错误造成的，如果不是那么就返回给客户端
		err = B.ServerProcessMes(&mess)

	}
}
