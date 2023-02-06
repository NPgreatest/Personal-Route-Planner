package router

import (
	"Personal-Route-Planner/controller"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	registerHomeRouters(engine)
	registerUserRouters(engine)
}

func registerHomeRouters(engine *gin.Engine) {
	homeRouter := controller.NewHomeRouter()
	siteRouter := controller.NewSiteRouter()
	homeGroup := engine.Group("/home")
	{
		homeGroup.GET("/getcomments", Decorate(siteRouter.GetComments))
		homeGroup.POST("/register", Decorate(homeRouter.RegisterUser))
		homeGroup.GET("/allsites", Decorate(siteRouter.FindAllSites))
		homeGroup.GET("/sitedetails", Decorate(siteRouter.FindSiteDetails))
	}
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
