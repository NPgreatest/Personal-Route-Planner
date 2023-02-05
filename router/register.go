package router

import (
	"Personal-Route-Planner/controller"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {

	registerUserRouters(engine)
}

func registerUserRouters(engine *gin.Engine) {
	userRouter := controller.NewUserRouter()
	engine.POST("/login", Decorate(userRouter.UserLogin))
	UserGroup := engine.Group("/user")
	UserGroup.Use(controller.LoginAuthenticationMiddleware())
	{
		UserGroup.POST("/comment", Decorate(userRouter.UserComment))
	}
}
