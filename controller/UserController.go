package controller

import (
	"Personal-Route-Planner/ChatGPT"
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"Personal-Route-Planner/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jung-kurt/gofpdf"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	userService   *service.UserService
	ratingService *service.RatingService
	siteService   *service.SiteService
}

func NewUserRouter() *UserController {
	return &UserController{userService: service.NewUserService(),
		ratingService: service.NewRatingService(),
		siteService:   service.NewSiteService()}
}

func (u *UserController) GetSummary(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("tagid"))
	if err != nil {
		return response.ResponseQueryFailed()
	}
	raw, err := u.userService.GetSummary(con)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("SimKai", "", "simkai.ttf")
	pdf.SetFont("SimKai", "", 14)
	pdf.MultiCell(0, 10, "团日活动总结", "", "C", false) // 添加标题
	pdf.MultiCell(0, 10, raw, "", "L", false)
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	ctx.Writer.Header().Set("Content-Type", "application/pdf")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=file.pdf")
	ctx.Writer.Write(buf.Bytes())
	return response.ResponseRegisterSuccess()
}

func (u *UserController) UserLogin(ctx *gin.Context) *response.Response {
	fmt.Println("vertifing password...")
	var user model.User
	err := ctx.ShouldBind(&user)
	user.Avatar = " "
	fmt.Println(user)
	if err != nil || user.Password == "" {
		return response.NewResponseOkND(response.LoginFailed)
	}
	resUser, err := u.userService.CheckUser(user.Name, user.Password)
	if err != nil {
		fmt.Println(user, err)
		return response.NewResponseOkND(response.LoginFailed)
	}
	token, err := utils.CreateToken(resUser.Name, time.Hour*24)
	if err != nil {
		fmt.Println(user, err)
		return response.NewResponseOkND(response.OperateFailed)
	}
	return response.NewResponseOk(response.LoginSuccess, token, resUser.Name)
}

func (u *UserController) UserComment(ctx *gin.Context) *response.Response {
	var comment model.Comment
	err := ctx.ShouldBind(&comment)
	if err != nil || comment.Sid == 0 || comment.Content == "" {
		//fmt.Println(comment, err)
		return response.ResponseCommentFailed()
	}
	comment.Cid = utils.GenerateId(1)
	comment.Name, _ = utils.VerifyToken(ctx.GetHeader("Authorization"))
	comment.Time = time.Now()
	err = u.userService.InsertComment(comment)
	if err != nil {
		//fmt.Println(comment, err)
		return response.ResponseCommentFailed()
	}
	return response.ResponseCommentSuccess()
}

func (u *UserController) UserRating(ctx *gin.Context) *response.Response {
	var rate model.Rating
	err := ctx.ShouldBindBodyWith(&rate, binding.JSON)
	if err != nil {
		fmt.Println(err)
		return response.ResponseRatingFailed()
	}
	name, _ := utils.VerifyToken(ctx.GetHeader("Authorization"))
	if err != nil {
		fmt.Println(err)
		return response.ResponseRatingFailed()
	}
	rate.Name = name
	fmt.Println(rate)
	err = u.ratingService.Rating(rate)
	if err != nil {
		fmt.Println(err)
		return response.ResponseRatingFailed()
	}
	return response.ResponseRatingSuccess()
}
func (u *UserController) UserGetRate(ctx *gin.Context) *response.Response {
	name, _ := utils.VerifyToken(ctx.GetHeader("Authorization"))
	var res string
	var err error
	if res, err = u.ratingService.CheckRating(name); err != nil {
		return response.ResponseQueryFailed()
	}
	var rate model.Rating
	err = json.Unmarshal([]byte(res), &rate)
	if err != nil {
		fmt.Println("反序列化错误 err=%v\n", err)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(rate)
}

func (u *UserController) GetRecommand(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("sid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	res, err := u.ratingService.GetRecommand("admin", con)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func (u *UserController) UserInfo(ctx *gin.Context) *response.Response {
	name, _ := utils.VerifyToken(ctx.GetHeader("Authorization"))
	res, err := u.userService.UserInfo(name)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func (u *UserController) UpdateUserInfo(ctx *gin.Context) *response.Response {
	var user model.User
	ctx.ShouldBind(&user)
	t, err := u.userService.UserInfo(user.Name)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	if user.Password == "" {
		user.Password = t.Password
	}
	if user.Avatar == "" {
		user.Avatar = t.Avatar
		u.userService.UpdateUserInfo(user)
		return response.ResponseOperateSuccess()
	}
	ava := strings.Split(t.Avatar, "/")
	err = os.Remove("./images/" + ava[4])
	if err != nil {
		fmt.Println(err)
	}
	user.Avatar, err = utils.WriteImages(user.Avatar)
	u.userService.UpdateUserInfo(user)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseRegisterSuccess()
}
func (u *UserController) SiteGPT(ctx *gin.Context) *response.Response {
	sid, err := strconv.Atoi(ctx.Query("sid"))
	str := ctx.Query("detail")
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	sname, err := u.siteService.SiteGPT(sid)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	res := ChatGPT.Callgpt("请介绍旅游景点" + sname + "\n" + str)
	return response.ResponseQuerySuccess(res)
}

func LoginAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			fmt.Println("未获得授权, ip:%s", ctx.Request.RemoteAddr)
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		if _, ok := utils.VerifyToken(token); !ok {
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
