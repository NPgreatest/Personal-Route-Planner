package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"github.com/gin-gonic/gin"
	"time"
)

type HomeController struct {
	homeService *service.HomeService
}

func NewHomeRouter() *HomeController {
	return &HomeController{homeService: service.NewHomeService()}
}

func (h *HomeController) GetComments(ctx *gin.Context) *response.Response {
	comments, err := h.homeService.GetComment()
	if err != nil {
		response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(comments)
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
		//fmt.Println(err, user)
		return response.ResponseRegisterFailed()
	}
	return response.ResponseRegisterSuccess()
}
