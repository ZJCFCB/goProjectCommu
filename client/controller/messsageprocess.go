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
		case util.MessageGroupInformType:
			HandMesGroupInform(mes.Data)
		case util.MessageSideInformType:
			HandMesSideInform(mes.Data)
		case util.MessageGroupResType:
			HandMesGroupRes(mes.Data)
		case util.MessageSideResType:
			HandMesSideRes(mes.Data)
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

func HandMesGroupInform(data string) {
	var mesgroup model.MesGroupInform
	err := json.Unmarshal([]byte(data), &mesgroup)
	if err != nil {
		fmt.Println(util.ERROR_UN_MARSHAL_FAILED)
	}

	if mesgroup.Errno == util.Success {
		fmt.Printf("收到了来自用户 %s 的群发消息，内容是 : %s \n", mesgroup.Name, mesgroup.Toall)
	} else {
		fmt.Println("获取群发消息失败 ", mesgroup.Message)
	}
}

func HandMesSideInform(data string) {
	var mesSide model.MesSideInform
	err := json.Unmarshal([]byte(data), &mesSide)
	if err != nil {
		fmt.Println(util.ERROR_UN_MARSHAL_FAILED)
	}

	if mesSide.Errno == util.Success {
		fmt.Printf("收到了来自用户 %s 的私信,它的id是 %d , 内容是 %s\n", mesSide.Namefrom, mesSide.Idfrom, mesSide.Side)
	} else {
		fmt.Println("获取私聊消息失败 ", mesSide.Message)
	}
}

func HandMesGroupRes(data string) {
	var mes model.MesGroupRes
	err := json.Unmarshal([]byte(data), &mes)
	if err != nil {
		fmt.Println(util.ERROR_UN_MARSHAL_FAILED)
	}

	if mes.Errno == util.Success {
		//fmt.Println("群发信息成功")
	} else {
		fmt.Println("群发信息失败 ", mes.Message)
	}
}

func HandMesSideRes(data string) {
	var mes model.MesSideRes
	err := json.Unmarshal([]byte(data), &mes)
	if err != nil {
		fmt.Println(util.ERROR_UN_MARSHAL_FAILED)
	}

	if mes.Errno == util.Success {
		//fmt.Printf("私发信息成功")
	} else {
		fmt.Println("私发信息失败 ", mes.Message)
	}
}
