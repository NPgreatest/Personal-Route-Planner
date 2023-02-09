package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
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
	DataArr := strings.Split(user.Avatar, ",")
	imgBase64 := DataArr[1]
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return response.ResponseOperateFailed()
	}
	timenow := time.Now().Unix()
	imgname := strconv.FormatInt(timenow, 10) + ".jpg"
	file, err := os.OpenFile("./images/"+imgname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return response.ResponseOperateFailed()
	}
	w := bufio.NewWriter(file)
	_, err = w.WriteString(string(imgs))
	if err != nil {
		return response.ResponseOperateFailed()
	}
	w.Flush()
	defer file.Close()
	user.Avatar = "http://localhost:8080/images/" + imgname
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
