package app

import (
	"github.com/Kutukobra/FinproKemjar_5/backend/app/middleware"
	"github.com/gin-gonic/gin"
)

func (a *App) Routes(router *gin.Engine) {
	// Serve static files
	router.Static("/static", "./app/static")

	// Page routes
	router.GET("/", a.pageHandler.DashboardPage)
	router.GET("/login", a.pageHandler.LoginPage)
	router.GET("/register", a.pageHandler.RegisterPage)
	router.GET("/dashboard", a.pageHandler.DashboardPage)
	router.GET("/change-password", a.pageHandler.ChangePasswordPage)
	router.GET("/profile", middleware.SessionAuth(), a.pageHandler.ProfilePage)

	// API routes (JSON responses)
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/:username", middleware.SessionAuth(), a.userHandler.GetUser)
			user.POST("/register", a.userHandler.RegisterUser)
			user.POST("/login", a.userHandler.LoginUser)
			user.PUT("/change-password", middleware.SessionAuth(), a.userHandler.ChangeUserPassword)
		}
	}

	// Form submission routes (redirect after success)
	router.POST("/login", a.userHandler.LoginUserForm)
	router.POST("/register", a.userHandler.RegisterUserForm)

}
