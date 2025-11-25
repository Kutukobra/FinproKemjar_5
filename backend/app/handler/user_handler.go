package handler

import (
	"net/http"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	serv *service.UserService
}

func NewUserHandler(serv *service.UserService) *UserHandler {
	return &UserHandler{serv: serv}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Query("username")

	userData, err := h.serv.GetUser(ctx, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	username, email, password := c.Query("username"), c.Query("email"), c.Query("password")

	userData, err := h.serv.RegisterUser(ctx, username, email, password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"data": userData,
		})
	}
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	ctx := c.Request.Context()

	username, password := c.Query("username"), c.Query("password")

	userData, err := h.serv.LoginUser(ctx, username, password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
	}
}

func (h *UserHandler) ChangeUserPassword(c *gin.Context) {
	ctx := c.Request.Context()

	username, password := c.Query("username"), c.Query("password")

	userData, err := h.serv.ChangeUserPassword(ctx, username, password)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
	}
}
