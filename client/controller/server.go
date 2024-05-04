package controller

import (
	"client/model"
	"client/util"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Userserve struct {
	Conn net.Conn
	Id   int
	Tf   *util.Transfer
}

//显示登录成功的界面

func (U *Userserve) ServerProcessMessage() {
	var key int
	var loop bool = false

	var Mg *MessageProcess = &MessageProcess{
		Channel: make(chan model.Message, 10),
	}

	go Mg.ReadFromService(U.Tf)
	go Mg.HandMessageFromService()
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("---------------------登录成功---------------------")
		fmt.Println("\t\t\t1.显示在线用户列表")
		fmt.Println("\t\t\t2.私聊发送消息")
		fmt.Println("\t\t\t3.群发消息")
		fmt.Println("\t\t\t4.退出登录")
		fmt.Printf("请选择(1-4) : ")
		fmt.Scanln(&key)
		switch key {
		case 1:
			err := U.OnlineList()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("群发消息")
		case 4:
			fmt.Println("退出登录")
			loop = util.Exit()
		default:
			fmt.Println("重新输入")
		}
		//无论如何，用户已经退出了。
		if loop {
			_, err := U.Exit()
			if err != nil {
				fmt.Println("用户退出失败 ", err)
			}
			break
		}
	}
}

func (U *Userserve) Exit() (bool, error) {

	var exitMes model.ExitMes
	exitMes.UserId = U.Id

	data, err := json.Marshal(exitMes)
	if err != nil {
		return false, util.ERROR_MARSHAL_FAILED
	}

	err = U.Tf.SendMessage(data, util.ExitType)

	if err != nil {
		return false, err
	}

	mes, err := U.Tf.ReadPkg()

	if err != nil {
		return false, err
	}

	var exitResMes model.ExitResMes

	err = json.Unmarshal([]byte(mes.Data), &exitResMes) // 对收到的数据反序列化

	if err != nil {
		return false, util.ERROR_UN_MARSHAL_FAILED
	}

	switch exitResMes.Errno {
	case util.Success:
		return true, nil
	case util.ExitFailed:
		return false, util.ERROR_EXIT_FAIL
	}
	return false, util.ERROR_UN_KNOW
}

func (U *Userserve) OnlineList() (err error) {

	var onlineMes model.OnlineListMes

	onlineMes.UserId = U.Id

	data, err := json.Marshal(onlineMes)

	if err != nil {
		return util.ERROR_MARSHAL_FAILED
	}

	err = U.Tf.SendMessage(data, util.OnlineListType)

	if err != nil {
		return err
	}

	return nil
}
