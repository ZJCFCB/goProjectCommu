package controller

import (
	"encoding/json"
	"server/model"
	"server/util"
)

func SendMessageGroup(toall string, id int, name string) error {
	var mes model.MesGroupRes
	mes.Errno = util.Success
	mes.Message = "成功"
	mes.Toall = toall
	mes.Idfrom = id
	mes.Name = name

	data, err := json.Marshal(mes)
	if err != nil {
		return util.ERROR_MARSHAL_FAILED
	}
	for k, v := range UserMgr.OnlineUser {
		if k == id {
			continue
		}
		v.Tf.SendMessage(data, util.MessageGroupResType)
	}
	return nil
}

func SendMessageSide(toside *model.MesSide) (err error) {
	var mes model.MesSideRes

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

		err = up.Tf.SendMessage(data, util.MessageSideResType)
		if err != nil {
			return err
		}
	}
	return nil
}
