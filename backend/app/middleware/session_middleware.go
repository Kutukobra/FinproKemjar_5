package middleware

import (
	"log"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		log.Println("Cookie: " + sessionToken)
		if err != nil || sessionToken == "" {
			// c.Redirect(http.StatusFound, "/login")
			c.Header("HX-Redirect", "/login")
			c.Abort()
		}
		claims, err := service.Validatetoken(sessionToken)
		if err != nil {
			// c.Redirect(http.StatusFound, "/login")
			c.Header("HX-Redirect", "/login")
			c.Abort()
		}
		c.Set("authUser", claims.Username)
		c.Next()
	}
}
