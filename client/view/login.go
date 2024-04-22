package view

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"goProjectCommu/util"
	"net"
)

func LoginCheck(id int, passwd string) (err error) {
	//建立连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net Dial failed")
		return nil
	}

	defer conn.Close()
	//准备发数据 message
	var mes util.Message
	mes.Type = util.LoginMesType

	//创建登录message
	var loginMes util.LoginMes
	loginMes.UserId = id
	loginMes.UserPwd = passwd

	//将这部分信息序列化，然后给data
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal failed", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal failed", err)
	}

	var pkglen uint32
	pkglen = uint32(len(data))
	var buffer [4]byte
	binary.BigEndian.PutUint32(buffer[:4], pkglen)

	_, err = conn.Write(buffer[:4])

	if err != nil {
		fmt.Println("conn.Write failed", err)
	}

	_, err = conn.Write(data)

	if err != nil {
		fmt.Println("conn.Write failed", err)
	}

	//处理返回的数据

	mes, err = util.ReadPkg(conn)

	if err != nil {
		fmt.Println("util.ReadPkg failed", err)
	}

	var loginResMes util.LoginRes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Errno == 200 {
		fmt.Println("登录成功")
		return nil
	} else {
		fmt.Println("登录失败")
		return errors.New("登录失败")
	}
}
