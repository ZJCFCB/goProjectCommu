package controller

import (
	"client/model"
	"client/util"
	"encoding/json"
	"net"
)

/*
这里主要是用户控制类
包括登录校验，用户注册，聊天等
*/
type UserProcess struct {
	Conn net.Conn
	Id   int
}

func (U *UserProcess) MakeConn(ip string) (err error) { //与传进来的ip建立链接（服务器）
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return util.ERROR_CONN_TO_SERVER_FAIL
	}
	U.Conn = conn
	return nil
}

func (U *UserProcess) LoginCheck(id int, passwd string) (isok bool, err error) {

	//model.LoginMes 封装登录信息，包括用户id、密码、用户名字
	var loginMes model.LoginMes
	loginMes.UserId = id
	loginMes.UserPwd = passwd

	//首先，将model.LoginMes 序列化，这部分是需要传输的内容
	//Marshal 序列化后的data 类型为 []byte
	data, err := json.Marshal(loginMes)
	if err != nil {
		return false, util.ERROR_MARSHAL_FAILED
	}

	// 用于控制收发数据
	tf := &util.Transfer{
		Conn: U.Conn,
	}

	err = tf.SendMessage(data, util.LoginMesType)

	if err != nil {
		return false, err
	}

	//处理返回的数据
	mes, err := tf.ReadPkg()

	if err != nil {
		return false, err
	}

	var loginResMes model.LoginRes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes) // 对收到的数据反序列化

	if err != nil {
		return false, util.ERROR_UN_MARSHAL_FAILED
	}

	switch loginResMes.Errno {
	case util.Success:
		return true, nil
	case util.NoRegistered:
		return false, util.ERROR_USER_NOTEXIT
	case util.PasswdIsWrong:
		return false, util.ERROR_PASSWD_RONG
	}
	return
}

func (U *UserProcess) Regist(id int, passwd, name string) (isok bool, err error) {

	//用于封装注册信息
	var registmes model.RegistMes
	registmes.UserId = id
	registmes.UserPwd = passwd
	registmes.UserName = name

	//序列化封装好的注册信息

	data, err := json.Marshal(registmes)

	if err != nil {
		return false, util.ERROR_MARSHAL_FAILED
	}

	//开始跟服务器端转发消息

	tf := &util.Transfer{Conn: U.Conn}

	err = tf.SendMessage(data, util.RegistMesType)
	if err != nil {
		return false, err
	}

	//然后等待读取服务器端返回的数据
	mes, err := tf.ReadPkg()

	if err != nil {
		return false, err
	}

	//然后对读取到的数据进行反序列化
	var registmesRes model.RegistRes

	err = json.Unmarshal([]byte(mes.Data), &registmesRes)
	if err != nil {
		return false, util.ERROR_READ_CONN_FAILED
	}
	switch registmesRes.Errno {
	case util.Success:
		return true, nil
	case util.UserHasExist:
		return false, util.ERROR_USER_IS_EXIST
	default:
		return false, util.ERROR_UN_KNOW
	}
}
