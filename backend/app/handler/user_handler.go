package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username."})
		return
	}

	userData, err := h.serv.GetUser(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	html := fmt.Sprintf(`
		<div class="data-display">
			<p><strong>ID:</strong> %s</p>
			<p><strong>Username:</strong> %s</p>
			<p><strong>Email:</strong> %s</p>
		</div>
	`, userData.ID, userData.Username, userData.Email)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
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

	_, err := h.serv.LoginUser(ctx, username, password)

	if err == nil {
		tokenString, err := service.GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		}

		c.SetCookie("session_token", tokenString, int(1*time.Hour), "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful.",
		})
		return
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

func (h *UserHandler) LoginUserForm(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.PostForm("username")
	password := c.PostForm("password")

	_, err := h.serv.LoginUser(ctx, username, password)

	if err == nil {
		c.Header("HX-Redirect", "/dashboard?username="+username)
		c.Status(http.StatusOK)
		return
	}

	// user not found
	if err == pgx.ErrNoRows {
		c.Writer.WriteString(`<div class="error-message">Wrong username or password</div>`)
		return
	}

	// wrong password
	if err == bcrypt.ErrMismatchedHashAndPassword {
		c.Writer.WriteString(`<div class="error-message">Wrong username or password</div>`)
		return
	}

	c.Writer.WriteString(`<div class="error-message">Server error occurred</div>`)
}

func (h *UserHandler) ChangeUserPassword(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.PostForm("username")
	currentPassword := c.PostForm("current_password")
	newPassword := c.PostForm("password")

	if username == "" || currentPassword == "" || newPassword == "" {
		c.Error(errors.New("empty username, current password or new password"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, current password and new password cannot be empty.",
		})
		return
	}

	_, err := h.serv.LoginUser(ctx, username, currentPassword)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword || err == pgx.ErrNoRows {
			c.Error(errors.New("incorrect current password"))
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Current password is incorrect.",
			})
			return
		}
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
		})
		return
	}

	_, err = h.serv.ChangeUserPassword(ctx, username, newPassword)

	if err == nil {
		c.Header("HX-Redirect", "/dashboard?username="+username)
		return
	}

	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error.",
	})
}

func (h *UserHandler) RegisterUserForm(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	log.Println(username + email + password)

	if username == "" || email == "" || password == "" {
		c.Writer.WriteString(`<div class="error-message">All fields are required</div>`)
		return
	}

	_, err := h.serv.RegisterUser(ctx, username, email, password)

	if err == nil {
		c.Header("HX-Redirect", "/login")
		c.Status(http.StatusOK)
		return
	}

	// duplicate username/email
	if pqErr, ok := err.(*pgconn.PgError); ok && pqErr.Code == "23505" {
		c.Writer.WriteString(`<div class="error-message">Username or Email already taken</div>`)
		return
	}

	c.Writer.WriteString(`<div class="error-message">Server error occurred</div>`)
}
