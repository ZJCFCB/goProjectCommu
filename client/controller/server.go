package controller

import (
	"client/model"
	"client/util"
	"encoding/json"
	"fmt"
	"net"
)

type Userserve struct {
	Conn net.Conn
	Id   int
}

//显示登录成功的界面

func (U *Userserve) ServerProcessMessage() {
	var key int
	var loop bool = false

	for {
		fmt.Println("---------------------登录成功---------------------")
		fmt.Println("\t\t\t1.显示在线用户列表")
		fmt.Println("\t\t\t2.发送消息")
		fmt.Println("\t\t\t3.显示登录信息列表")
		fmt.Println("\t\t\t4.退出登录")
		fmt.Printf("请选择(1-4) : ")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
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
	var mes model.Message
	mes.Type = util.ExitType

	var exitMes model.ExitMes
	exitMes.UserId = U.Id

	data, err := json.Marshal(exitMes)
	if err != nil {
		return false, util.ERROR_MARSHAL_FAILED
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)

	if err != nil {
		return false, util.ERROR_MARSHAL_FAILED
	}

	tf := util.Transfer{
		Conn: U.Conn,
	}

	err = tf.WritePkg(data)

	if err != nil {
		return false, err
	}

	mes, err = tf.ReadPkg()

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
