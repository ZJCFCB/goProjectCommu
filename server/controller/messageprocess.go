package controller

import (
	"encoding/json"
	"server/model"
	"server/util"
)

func SendMessageGroup(toall string, id int, name string) error {
	var mes model.MesGroupInform
	mes.Errno = util.Success
	mes.Message = "成功"
	mes.Toall = toall
	mes.Idfrom = id
	mes.Name = name

	data, err := json.Marshal(mes)
	if err != nil {
		return util.ERROR_MARSHAL_FAILED
	}
	//遍历在线列表，群发
	for k, v := range UserMgr.OnlineUser {
		if k == id {
			continue
		}
		err = v.Tf.SendMessage(data, util.MessageGroupInformType)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendMessageSide(toside *model.MesSide) (err error) {
	var mes model.MesSideInform

	up, ok := UserMgr.OnlineUser[toside.ToId]

	if !ok {
		//此用户不在线
		return util.ERROR_USER_NOT_ONLINE
	} else {
		mes.Errno = util.Success
		mes.Message = "成功"
		mes.Idfrom = toside.MyId
		mes.Namefrom = toside.MyName
		mes.Side = toside.Side

		data, err := json.Marshal(mes)

		if err != nil {
			return util.ERROR_MARSHAL_FAILED
		}

		err = up.Tf.SendMessage(data, util.MessageSideInformType)
		if err != nil {
			return err
		}
	}
	return nil
}
