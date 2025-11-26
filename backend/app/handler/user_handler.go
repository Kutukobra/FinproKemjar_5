package handler

import (
	"errors"
	"net/http"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
		return
	}

	if err == pgx.ErrNoRows {
		c.Error(errors.New("user not found"))
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found.",
		})
		return
	}

	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error.",
	})
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	username, email, password := c.Query("username"), c.Query("email"), c.Query("password")

	userData, err := h.serv.RegisterUser(ctx, username, email, password)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"data": userData,
		})
		return
	}

	if pqErr, ok := err.(*pgconn.PgError); ok && pqErr.Code == "23505" {
		c.Error(errors.New("username or email taken"))
		c.JSON(http.StatusConflict, gin.H{
			"error": "Username or Email taken.",
		})
		return
	}

	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error.",
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	ctx := c.Request.Context()

	username, password := c.Query("username"), c.Query("password")

	userData, err := h.serv.LoginUser(ctx, username, password)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
	}

	if err == bcrypt.ErrMismatchedHashAndPassword {
		c.Error(errors.New("wrong email or password"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong email or password.",
		})
		return
	}

	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error.",
	})
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

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
		return
	}

	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error.",
	})
}
