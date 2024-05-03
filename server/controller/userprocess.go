package controller

import (
	"encoding/json"
	"net"
	"server/model"
	"server/service"
	"server/util"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
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
		//登录成功后，把在线用户的id和userprecess放入在线列表中
		U.UserId = loginMessage.UserId
		UserMgr.AddOnlineUser(U)

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

func (U *UserProcess) HandRegist(mes *model.Message) (err error) {

	//返回消息
	var registRes model.Message
	registRes.Type = util.RegistResMesType

	//先从mes中取出信息并反序列化
	var registmes model.RegistMes

	err = json.Unmarshal([]byte(mes.Data), &registmes)

	if err != nil {
		return util.ERROR_UN_MARSHAL_FAILED
	}

	// 开始注册
	var userservice service.UserService = service.UserService{}
	isSuccess, err := userservice.Regist(registmes.UserId, registmes.UserPwd, registmes.UserName)

	//处理返回结果
	var registResModel model.RegistRes
	if isSuccess {
		registResModel.Errno = util.Success
		registResModel.Message = "注册成功"
	} else {
		if err == util.ERROR_USER_HAS_EXIT {
			registResModel.Errno = util.UserHasExist
			registResModel.Message = "注册失败，用户已经存在"
		} else {
			registResModel.Errno = util.SERVICE_HAS_WRONG
			registResModel.Message = "服务器内部错误"
		}
	}

	data, err := json.Marshal(registResModel)

	if err != nil {
		return util.ERROR_MARSHAL_FAILED
	}

	registRes.Data = string(data)

	data, err = json.Marshal(registRes)

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

func (U *UserProcess) HandExit(mes *model.Message) (err error) {
	var mesRes model.Message
	mesRes.Type = util.ExitType

	var exitRes model.ExitResMes
	var exitMes model.ExitMes

	err = json.Unmarshal([]byte(mes.Data), &exitMes)
	if err != nil {
		return util.ERROR_UN_MARSHAL_FAILED
	}

	UserMgr.DelOnlineUser(exitMes.UserId)

	exitRes.Errno = util.Success
	exitRes.Message = "用户成功退出"

	data, err := json.Marshal(exitRes)

	if err != nil {
		return err
	}

	mesRes.Data = string(data)

	data, err = json.Marshal(mesRes)

	if err != nil {
		return err
	}

	// 发送data
	var tf *util.Transfer = &util.Transfer{
		Conn: U.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return err
	}
	return util.ERROR_EXIT_SUCCESS

}
