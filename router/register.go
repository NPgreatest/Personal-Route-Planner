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
	tagRouter := controller.NewTagRouter()
	homeGroup := engine.Group("/api/home")
	{
		homeGroup.GET("/getcomments", Decorate(siteRouter.GetComments))
		homeGroup.POST("/register", Decorate(homeRouter.RegisterUser))
		homeGroup.GET("/allsites", Decorate(siteRouter.FindAllSites))
		homeGroup.GET("/sitedetails", Decorate(siteRouter.FindSiteDetails))
		homeGroup.GET("/alltags", Decorate(tagRouter.FindAllTags))
		homeGroup.GET("/sitesbytags", Decorate(tagRouter.FindSitesByTags))
	}
}

func registerUserRouters(engine *gin.Engine) {
	userRouter := controller.NewUserRouter()
	engine.POST("/api/login", Decorate(userRouter.UserLogin))
	UserGroup := engine.Group("/api/user")
	UserGroup.Use(controller.LoginAuthenticationMiddleware())
	{
		UserGroup.POST("/comment", Decorate(userRouter.UserComment))
		UserGroup.POST("/rate", Decorate(userRouter.UserRating))
		UserGroup.GET("/getrate", Decorate(userRouter.UserGetRate))
	}
}
