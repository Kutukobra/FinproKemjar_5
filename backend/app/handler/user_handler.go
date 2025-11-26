package handler

import (
	"errors"
	"net/http"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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

	username := c.Param("username")

	if username == "" {
		c.Error(errors.New("invalid username"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username.",
		})
		return
	}

	userData, err := h.serv.GetUser(ctx, username)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
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

	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			c.Error(pqErr)
			c.JSON(http.StatusConflict, gin.H{
				"error": "Username or Email already taken.",
			})
		}
	}

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
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

	if err == bcrypt.ErrMismatchedHashAndPassword {
		c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong email or password.",
		})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
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

	if username == "" || password == "" {
		c.Error(errors.New("empty username or password"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or password cannot be empty.",
		})
		return
	}

	userData, err := h.serv.ChangeUserPassword(ctx, username, password)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
	}
}
