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
