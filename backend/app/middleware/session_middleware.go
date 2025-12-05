package middleware

import (
	"net/http"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil || sessionToken == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}
		_, err = service.Validatetoken(sessionToken)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}
		c.Next()
	}
}
