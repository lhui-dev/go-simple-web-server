// 用户服务层
// author 李少辉

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/leedev/go-simple-web-server/api/dao"
	"github.com/leedev/go-simple-web-server/internal/common/result"
)

// UserService 用户服务层接口定义
type UserService interface {
	QueryUserList(ctx *gin.Context)
}

type UserServiceImpl struct {
	userDao dao.UserDao
}

func (us *UserServiceImpl) QueryUserList(ctx *gin.Context) {
	userVoList := us.userDao.QueryUserVoList()
	data := map[string]interface{}{"user": userVoList}
	result.Success(ctx, data)
}

func NewUserServiceImpl(userDao dao.UserDao) UserService {
	return &UserServiceImpl{userDao: userDao}
}
