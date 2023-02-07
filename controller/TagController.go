package controller

import (
	"Personal-Route-Planner/db/service"
	"Personal-Route-Planner/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TagController struct {
	tagService *service.TagService
}

func NewTagRouter() *TagController {
	return &TagController{tagService: service.NewTagService()}
}

func (t *TagController) FindAllTags(ctx *gin.Context) *response.Response {
	tags, err := t.tagService.FindAllTags()
	if err != nil {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(tags)
}

func (t *TagController) FindSitesByTags(ctx *gin.Context) *response.Response {
	con, err := strconv.Atoi(ctx.Query("siteid"))
	if err != nil {
		fmt.Println("siteid wrong")
		return response.ResponseQueryFailed()
	}
	sites, err := t.tagService.FindSitesByTags(con)
	if err != nil {
		fmt.Println("finding error")
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(sites)
}
