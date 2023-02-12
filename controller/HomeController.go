package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"Personal-Route-Planner/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type HomeController struct {
	homeService *service.HomeService
	userService *service.UserService
}

func NewHomeRouter() *HomeController {
	return &HomeController{homeService: service.NewHomeService(), userService: service.NewUserService()}
}

func (h *HomeController) FindName(name string) int {
	can, err := h.userService.FindName(name)
	if err != nil {
		return -1
	}
	if can == false {
		return 0
	}
	return 1
}

func (h *HomeController) RegisterUser(ctx *gin.Context) *response.Response {
	var user model.User
	ctx.ShouldBind(&user)
	same := h.FindName(user.Name)
	fmt.Println(same)
	switch same {
	case -1:
		return response.ResponseQueryFailed()
	case 0:
		return response.ResponseRegisterSameName()
	}
	var err error
	user.Avatar, err = utils.WriteImages(user.Avatar)
	user.CreateTime = time.Now()
	if user.Avatar == "" || user.Email == "" || user.Name == "" {
		return response.ResponseRegisterFailed()
	}
	err = h.homeService.RegisterUser(user)
	if err != nil {
		fmt.Println(err, user)
		return response.ResponseRegisterFailed()
	}
	return response.ResponseRegisterSuccess()
}
