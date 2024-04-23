package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"server/model"
)

type Transfer struct {
	Conn   net.Conn
	Buffer [8096]byte
}

func (T *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度给客户端

	var pkglen uint32
	pkglen = uint32(len(data))
	binary.BigEndian.PutUint32(T.Buffer[:4], pkglen)

	_, err = T.Conn.Write(T.Buffer[:4])

	if err != nil {
		fmt.Println("conn.Write failed", err)
	}

	_, err = T.Conn.Write(data)

	if err != nil {
		fmt.Println("conn.Write failed", err)
	}
	return
}

func (T *Transfer) ReadPkg() (mes model.Message, err error) {
	//读取客户端发来的数据
	fmt.Println("waitting reading")
	_, err = T.Conn.Read(T.Buffer[:4])
	if err != nil {
		fmt.Println("conn.Read failed")
		return
	}
	var pkglen uint32 = binary.BigEndian.Uint32(T.Buffer[:4])

	n, err := T.Conn.Read(T.Buffer[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Read failed")
		return
	}

	//反序列化package
	err = json.Unmarshal(T.Buffer[:pkglen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal failed")
		return
	}
	return
}
