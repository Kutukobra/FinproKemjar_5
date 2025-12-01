package handler

import (
	"net/http"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/template/pages"
	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	userService *service.UserService
}

func NewPageHandler(userService *service.UserService) *PageHandler {
	return &PageHandler{
		userService: userService,
	}
}

func (h *PageHandler) LoginPage(c *gin.Context) {
	component := pages.Login()
	component.Render(c.Request.Context(), c.Writer)
}

func (h *PageHandler) RegisterPage(c *gin.Context) {
	component := pages.Register()
	component.Render(c.Request.Context(), c.Writer)
}

func (h *PageHandler) DashboardPage(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	component := pages.Dashboard(username)
	component.Render(c.Request.Context(), c.Writer)
}

func (h *PageHandler) ChangePasswordPage(c *gin.Context) {
	component := pages.ChangePassword()
	component.Render(c.Request.Context(), c.Writer)
}

func (h *PageHandler) ProfilePage(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	component := pages.Profile(username)
	component.Render(c.Request.Context(), c.Writer)
}

func (h *PageHandler) HomePage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
}
