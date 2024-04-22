package controller

import (
	"encoding/json"
	"fmt"
	"net"
	"server/util"
)

type UserProcess struct {
	Conn net.Conn
}

func (U *UserProcess) HandLogin(mes *util.Message) (err error) {
	//先从message中取出data，并反序列化成login
	var loginMessage util.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMessage)
	if err != nil {
		fmt.Println("unmarshal failed")
		return
	}

	var resMessage util.Message
	resMessage.Type = util.LoginResMesType

	var loginres util.LoginRes

	if loginMessage.UserId == 100 && loginMessage.UserPwd == "123456" { //合法
		loginres.Errno = util.Success
		loginres.Message = "登录成功"
	} else { //不合法用户
		loginres.Errno = util.NoRegistered
		loginres.Message = "用户不存在"
	}
	data, err := json.Marshal(loginres)
	if err != nil {
		fmt.Println("xuliehuashibai ")
	}

	resMessage.Data = string(data)

	data, err = json.Marshal(resMessage)
	if err != nil {
		fmt.Println("xuliehuashibai ")
	}

	// 发送data
	var tf *util.Transfer = &util.Transfer{
		Conn: U.Conn,
	}
	tf.WritePkg(data)
	return
}
