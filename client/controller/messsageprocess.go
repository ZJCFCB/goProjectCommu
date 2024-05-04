package controller

import (
	"client/model"
	"client/util"
	"encoding/json"
	"fmt"
)

type MessageProcess struct {
	Channel chan model.Message
}

//这里用于管理client接受服务器的消息

func (MP *MessageProcess) ReadFromService(tf *util.Transfer) {
	for {
		data, err := tf.ReadPkg()
		if err != nil {
			break
		}
		MP.Channel <- data
	}
}

func (MP *MessageProcess) HandMessageFromService() {
	//用于处理管道内传过来的信息
	for {
		mes := <-MP.Channel
		switch mes.Type {
		case util.OnlineListType:
			HandLoginResMes(mes.Data)
		}

	}
}

func HandLoginResMes(data string) {
	var onlineRes model.OnlineListRes
	err := json.Unmarshal([]byte(data), &onlineRes)
	if err != nil {
		fmt.Println(util.ERROR_UN_MARSHAL_FAILED)
	}
	if onlineRes.Errno == util.Success {
		fmt.Println("在线用户列表是 ", onlineRes.OnlineList)
	} else {
		fmt.Println("获取用户列表失败，", onlineRes.Message)
	}
}
