package util

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"server/model"
)

/*
这里主要用于控制与服务端之间的通信
包括对conn写数据和读数据
在写数据之前，会先写入数据的长度，避免丢包
*/
type Transfer struct {
	Conn   net.Conn
	Buffer [8096]byte
}

/*
向Conn写入数据，注意参数类型是[]byte
*/
func (T *Transfer) WritePkg(data []byte) (err error) {

	//先发送长度给客户端
	var pkglen uint32
	pkglen = uint32(len(data))
	binary.BigEndian.PutUint32(T.Buffer[:4], pkglen)

	_, err = T.Conn.Write(T.Buffer[:4])

	if err != nil {
		return ERROR_WRITE_CONN_FAILED
	}

	_, err = T.Conn.Write(data)

	if err != nil {
		return ERROR_WRITE_CONN_FAILED
	}
	return nil
}

func (T *Transfer) ReadPkg() (mes model.Message, err error) {
	//读取客户端发来的数据,先读取到的是消息的长度
	_, err = T.Conn.Read(T.Buffer[:4])
	if err != nil {
		return mes, ERROR_READ_CONN_FAILED
	}

	//读取发送过来的消息
	var pkglen uint32 = binary.BigEndian.Uint32(T.Buffer[:4])
	n, err := T.Conn.Read(T.Buffer[:pkglen])
	if n != int(pkglen) || err != nil {
		return mes, ERROR_READ_CONN_FAILED
	}

	//反序列化package
	err = json.Unmarshal(T.Buffer[:pkglen], &mes)
	if err != nil {
		return mes, ERROR_UN_MARSHAL_FAILED
	}
	return mes, nil
}
