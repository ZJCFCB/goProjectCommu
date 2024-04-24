package service

import (
	"server/model"
	"server/model/dao"
	"server/util"
)

type UserService struct {
}

func (U *UserService) Login(id int, pwd string) (user model.User, err error) {

	user, err = dao.MyUserDao.GetUserById(id)
	if err != nil {
		return user, err
	}

	//校验密码
	if user.UserPwd == pwd { //密码正确,登录成功
		return user, nil
	} else {
		return user, util.ERROR_PASSWD_RONG
	}
}
