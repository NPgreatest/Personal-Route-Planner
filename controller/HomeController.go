package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type HomeController struct {
	homeService *service.HomeService
}

func NewHomeRouter() *HomeController {
	return &HomeController{homeService: service.NewHomeService()}
}

func (h *HomeController) RegisterUser(ctx *gin.Context) *response.Response {
	var user model.User
	ctx.ShouldBind(&user)
	user.CreateTime = time.Now()
	//fmt.Println(user)
	if user.Avatar == "" || user.Email == "" || user.Name == "" {
		return response.ResponseRegisterFailed()
	}
	err := h.homeService.RegisterUser(user)
	if err != nil {
		fmt.Println(err, user)
		return response.ResponseRegisterFailed()
	}
	return response.ResponseRegisterSuccess()
}
