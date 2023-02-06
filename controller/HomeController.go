package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type HomeController struct {
	homeService *service.HomeService
}

func NewHomeRouter() *HomeController {
	return &HomeController{homeService: service.NewHomeService()}
}

func (h *HomeController) GetComments(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("siteid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	comments, err := h.homeService.GetComment(con)
	if err != nil {
		response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(comments)
}

func (h *HomeController) RegisterUser(ctx *gin.Context) *response.Response {
	var user model.User
	ctx.ShouldBind(&user)
	user.CreateTime = time.Now()
	fmt.Println(user)
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

func (h *HomeController) FindAllSites(ctx *gin.Context) *response.Response {
	res, err := h.homeService.FindAllSites()
	if err != nil {
		fmt.Println(err)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func (h *HomeController) FindSiteDetails(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("siteid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	res, err := h.homeService.FindSiteDetails(con)
	if err != nil {
		fmt.Println("could not find site details", res)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}
