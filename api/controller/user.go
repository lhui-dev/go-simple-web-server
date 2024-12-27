// 用户控制层
// author 李少辉

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leedev/go-simple-web-server/api/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// QueryUserList godoc
//
// @Summary 	查询用户列表接口
// @Description 查询用户列表信息
// @Tags 		User
// @Produce 	json
// @Success 	200 		{object} 	result.Result
// @Failure		500			{object}	result.Result
// @Router 		/user/list 	[get]
func (controller *UserController) QueryUserList(ctx *gin.Context) {
	controller.userService.QueryUserList(ctx)
}
