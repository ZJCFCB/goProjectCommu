package controller

import (
	"client/model"
	"client/util"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (U *UserProcess) MakeConn(ip string) (err error) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return err
	}
	U.Conn = conn
	return nil
}

func (U *UserProcess) LoginCheck(id int, passwd string) (isok bool, err error) {
	//准备发数据 message
	var mes model.Message
	mes.Type = util.LoginMesType

	//创建登录message
	var loginMes model.LoginMes
	loginMes.UserId = id
	loginMes.UserPwd = passwd

	//将这部分信息序列化，然后给data
	data, err := json.Marshal(loginMes) //Marshal 序列化后的data 类型为 []byte
	if err != nil {
		return false, err
	}
	mes.Data = string(data) //把登录信息放在请求结构体的data部分

	data, err = json.Marshal(mes)
	if err != nil {
		return false, err
	}

	// 用于控制收发数据
	tf := &util.Transfer{
		Conn: U.Conn,
	}

	err = tf.WritePkg(data) //发送数据

	if err != nil {
		fmt.Println("客户端发送数据失败")
		return false, err
	}

	//处理返回的数据

	mes, err = tf.ReadPkg()

	if err != nil {
		return false, err
	}

	var loginResMes model.LoginRes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes) // 对收到的数据反序列化

	if err != nil {
		return false, err
	}

	if loginResMes.Errno == 200 {
		fmt.Println("登录成功")
		for {
			go ServerProcessMessage(U.Conn)
			ShowLoginMenu()
		}
	} else {
		fmt.Println("账户名或密码错误")
	}
	return
}
