package dao

import (
	"encoding/json"
	"server/model"

	"github.com/garyburd/redigo/redis"
)

//完成对User结构体对数据库的各种操作

//我们希望在启动的时候就完成UserDao的初始化，做成全局变量

var MyUserDao *UserDao

type UserDao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *UserDao {
	return &UserDao{
		Pool: pool,
	}
}

func InitUserDao() {
	MyUserDao = NewUserDao(RedisPool)
}

// 根据用户ID，返回实例
func (U *UserDao) getUserById(conn redis.Conn, id int) (user model.User, err error) {
	res, err := redis.String(conn.Do("HGet", "user", id))
	if err != nil {
		if err == redis.ErrNil { // 表示不存在
			err = model.ERROR_USER_NOTEXIT
			return
		}
		return
	}

	//需要把res反序列成user实例

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return
	}
	return
}

func (U *UserDao) Login(id int, pwd string) (user model.User, err error) {
	conn := U.Pool.Get() // 从连接池中取出一个链接
	defer conn.Close()

	user, err = U.getUserById(conn, id)
	if err != nil {
		return
	}

	//校验密码

	if user.UserPwd == pwd {
		return user, nil
	} else {
		return user, model.ERROR_PASSWD_RONG
	}
}
