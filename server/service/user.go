package service

import (
	"encoding/json"
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

func (U *UserService) Regist(id int, pwd, name string) (isok bool, err error) {

	isExit, err := dao.MyUserDao.IdIsExist(id)

	if isExit == false && err == nil { // 开始注册工作吧
		var user model.User
		user.UserId = int64(id)
		user.UserPwd = pwd
		user.UserName = name
		data, err := json.Marshal(user)
		if err != nil {
			return false, util.ERROR_MARSHAL_FAILED
		}

		stringDdata := string(data)

		err = dao.MyUserDao.SetInfromToRedis(id, stringDdata)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	if isExit == true {
		return false, util.ERROR_USER_HAS_EXIT
	}

	return false, err
}
