package controller

import (
	"encoding/json"
	"net"
	"server/model"
	"server/service"
	"server/util"
)

type UserProcess struct {
	Conn net.Conn
}

func (U *UserProcess) HandLogin(mes *model.Message) (err error) {
	//先从message中取出data，并反序列化成login
	var loginMessage model.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMessage)
	if err != nil {
		return util.ERROR_UN_MARSHAL_FAILED
	}

	//返回消息
	var resMessage model.Message
	resMessage.Type = util.LoginResMesType

	//返回登录信息
	var loginres model.LoginRes
	var userservice service.UserService = service.UserService{}

	//进行登录校验
	_, err = userservice.Login(loginMessage.UserId, loginMessage.UserPwd)

	// 根据error 决定返回的状态码是多少
	switch err {
	case nil:
		loginres.Errno = util.Success
		loginres.Message = "登录成功"
	case util.ERROR_USER_NOTEXIT:
		loginres.Errno = util.NoRegistered
		loginres.Message = "用户不存在"
	case util.ERROR_PASSWD_RONG:
		loginres.Errno = util.PasswdIsWrong
		loginres.Message = "用户名或密码错误"
	default:
		loginres.Errno = util.SERVICE_HAS_WRONG
		loginres.Message = "服务器内部发生错误"
	}

	data, err := json.Marshal(loginres)
	if err != nil {
		return util.ERROR_MARSHAL_FAILED
	}

	resMessage.Data = string(data)

	data, err = json.Marshal(resMessage)
	if err != nil {
		return util.ERROR_MARSHAL_FAILED
	}

	// 发送data
	var tf *util.Transfer = &util.Transfer{
		Conn: U.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return err
	}
	return nil
}
