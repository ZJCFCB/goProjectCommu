package dao

import (
	"encoding/json"
	"server/model"
	"server/util"

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

// 根据用户ID，返回redis中的用户信息
func (U *UserDao) GetUserById(id int) (user model.User, err error) {

	// 从连接池中取出一个链接
	conn := U.Pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("HGet", "user", id))
	if err != nil {

		if err == redis.ErrNil {
			return user, util.ERROR_USER_NOTEXIT
		}
		return user, err
	}

	//需要把res反序列成user实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return user, util.ERROR_UN_MARSHAL_FAILED
	}
	return user, nil
}

func (U *UserDao) IdIsExist(id int) (ok bool, err error) {

	conn := U.Pool.Get()
	defer conn.Close()

	_, err = redis.String(conn.Do("HGet", "user", id))

	if err != nil {
		if err == redis.ErrNil {
			return false, nil
		} else {
			return false, util.ERROR_READ_FROM_REDIS_FAILED //有报错的话，直接抛出去
		}
	}
	return true, nil //查询成功了
}

func (U *UserDao) SetInfromToRedis(id int, data string) (err error) {

	conn := U.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("HSet", "user", id, data)

	if err != nil {
		return util.ERROR_PUSH_TO_REDIS_FAILED
	}
	return nil
}
