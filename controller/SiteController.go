package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SiteController struct {
	siteService *service.SiteService
}

func NewSiteRouter() *SiteController {
	return &SiteController{siteService: service.NewSiteService()}
}

func (s *SiteController) GetComments(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("siteid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	comments, err := s.siteService.GetComment(con)
	if err != nil {
		response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(comments)
}

func (s *SiteController) FindAllSites(ctx *gin.Context) *response.Response {
	res, err := s.siteService.FindAllSites()
	if err != nil {
		fmt.Println(err)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func (s *SiteController) FindSiteDetails(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("siteid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	res, err := s.siteService.FindSiteDetails(con)
	if err != nil {
		fmt.Println("could not find site details", res)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}

func (s *SiteController) FindCommentsAvatars(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("siteid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	res, err := s.siteService.FindCommentsAvatars(con)
	if err != nil {
		fmt.Println("could not find site details", res)
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(res)
}
