package controller

import (
	"fmt"
	"server/util"
)

//用于用户管理
//凡是登录上的用户，都会把id和UserPreocess放入OnlineUser中
//主要是以为UserProcess中包括了用户与服务器的连接conn

var UserMgr *userMgr

type userMgr struct {
	OnlineUser map[int]*UserProcess
}

func init() {
	UserMgr = &userMgr{
		OnlineUser: make(map[int]*UserProcess, 64),
	}
}

func (U *userMgr) AddOnlineUser(up *UserProcess) {
	U.OnlineUser[up.UserId] = up
}

func (U *userMgr) DelOnlineUser(userid int) {
	delete(U.OnlineUser, userid)
}

func (U *userMgr) GetAllOnlineUser() map[int]*UserProcess {
	return U.OnlineUser
}

func (U *userMgr) GetOnlineUserById(id int) (up *UserProcess, err error) {
	up, ok := U.OnlineUser[id]

	if !ok { // 当前不在线
		return up, util.ERROR_USER_NOT_ONLINE
	}

	return up, nil
}

func (U *userMgr) PrintUser() {
	for k, _ := range U.OnlineUser {
		fmt.Println("the online user is ", k)
	}
}
