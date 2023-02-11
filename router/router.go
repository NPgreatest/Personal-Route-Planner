package router

import (
	"Personal-Route-Planner/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler func(ctx *gin.Context) *response.Response

// 装饰器
func Decorate(h Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "True")
		//放行索引options
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		ctx.Next()
		r := h(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		response.PutResponse(r)
	}
}
