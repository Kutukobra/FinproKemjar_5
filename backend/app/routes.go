package app

import (
	"github.com/gin-gonic/gin"
)

func (a *App) Routes(router *gin.Engine) {

	api := router.Group("/api")

	user := api.Group("/user")
	{
		user.GET("/:username", a.userHandler.GetUser)
		user.POST("/register", a.userHandler.RegisterUser)
		user.POST("/login", a.userHandler.LoginUser)
		user.PUT("/change-password", a.userHandler.ChangeUserPassword)
	}
}
