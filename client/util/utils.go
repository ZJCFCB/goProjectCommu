package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func ReadPkg(conn net.Conn) (mes Message, err error) {
	//读取客户端发来的数据
	var buffer []byte = make([]byte, 8096)
	fmt.Println("waitting reading")
	_, err = conn.Read(buffer[:4])
	if err != nil {
		fmt.Println("conn.Read failed")
		return
	}
	var pkglen uint32 = binary.BigEndian.Uint32(buffer[:4])

	n, err := conn.Read(buffer[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Read failed")
		return
	}

	//反序列化package
	err = json.Unmarshal(buffer[:pkglen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal failed")
		return
	}
	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {
	//先发送长度给客户端

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
	return
}
