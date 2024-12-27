package dao

import (
	"github.com/leedev/go-simple-web-server/api/model/entity"
	"github.com/leedev/go-simple-web-server/api/model/vo"
)

// UserDao 用户数据层接口定义
type UserDao interface {
	// QueryUserByUsername 根据用户名查询用户信息
	QueryUserByUsername(username string)

	// QueryUserVoList 查询用户信息列表
	QueryUserVoList() []vo.UserVo
}

type UserDaoImpl struct {
}

func (userDao *UserDaoImpl) QueryUserVoList() []vo.UserVo {
	users := []entity.User{
		{
			UserId:   "1",
			UserName: "lee",
			NickName: "aaa",
		},
		{
			UserId:   "2",
			UserName: "bob",
			NickName: "BBB",
		},
	}
	var userVoList []vo.UserVo
	for _, user := range users {
		userVo := vo.UserVo{
			UserId:   user.UserId,
			UserName: user.UserName,
			NickName: user.NickName,
		}
		userVoList = append(userVoList, userVo)
	}
	return userVoList
}

func NewUserDaoImpl() UserDao {
	return &UserDaoImpl{}
}

func (userDao *UserDaoImpl) QueryUserByUsername(username string) {
	//TODO implement me
	panic("implement me")
}
