package controller

import (
	"Personal-Route-Planner/Node"
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/model"
	"Personal-Route-Planner/response"
	"Personal-Route-Planner/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	userService     *service.UserService
	ratingService   *service.RatingService
	siteService     *service.SiteService
	activityService *service.ActivityService
	tagService      *service.TagService
}

func NewUserRouter() *UserController {
	return &UserController{userService: service.NewUserService(),
		ratingService:   service.NewRatingService(),
		siteService:     service.NewSiteService(),
		activityService: service.NewActivityService(),
		tagService:      service.NewTagService(),
	}
}

func (u *UserController) GetSummary(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("tagid"))
	t1 := ctx.Query("sites")
	t2 := ctx.Query("activities")
	tag, err := u.tagService.FindTagName(con)
	raw, err := u.userService.GetSummary(rand.Int()%5 + 1)
	if err != nil {
		return response.ResponseQueryFailed()
	}
	var sites []string
	for _, s := range strings.Split(t1[1:len(t1)-1], ",") {
		if s != "\"\"" {
			sites = append(sites, s)
		}
	}
	activity := make([]string, 0)
	for _, s := range strings.Split(t2[1:len(t2)-1], ",") {
		if s != "\"\"" {
			activity = append(activity, s)
		}
	}
	q := "请进行活动主题为" + tag + "活动总结：我们从" + sites[0] + "出发，分别在"
	for index, value := range sites {
		if index == 0 {
			continue
		}
		q += fmt.Sprintf(value + "进行了活动" + activity[index-1] + "，")
	}

	res, err := Node.QueryNodeChatGPT(q)
	if err != nil {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(raw + res)
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
	history := ctx.Query("history")
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	if err != nil {
		return response.ResponseQueryFailed()
	}
	res, err := Node.QueryNodeQuery(str, sid, history)
	if err != nil {
		fmt.Println(err)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func (u *UserController) GetActivity(ctx *gin.Context) *response.Response {
	sid, err := strconv.Atoi(ctx.Query("sid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	res, err := u.activityService.GetActivity(sid)
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	type t struct {
		name string `json:"name"`
	}
	var temp []t
	for _, v := range res {
		temp = append(temp, t{name: v})
	}
	return response.ResponseQuerySuccess(res)
}

func (u *UserController) GetActivityDescription(ctx *gin.Context) *response.Response {
	sname := ctx.Query("sname")
	description := ctx.Query("description")
	res, err := Node.QueryNodeChatGPT("请模拟官方的身份介绍在景点" + sname + "举办的" + description + "活动。")
	if err != nil {
		fmt.Println("GetActivityDescription wrong")
		return response.ResponseQueryFailed()
	}
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
