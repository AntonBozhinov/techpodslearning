package middleware

import (
	"github.com/AntonBozhinov/techpodslearn/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Protect(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("techpods")
		if err != nil {
			http.Redirect(c.Writer, c.Request, "/auth/google/login", http.StatusTemporaryRedirect)
			return
		}
		token, err := jwt.ParseWithClaims(cookie, &user.TokenClaims{}, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			http.Redirect(c.Writer, c.Request, "/auth/google/login", http.StatusTemporaryRedirect)
			return
		}
		if claims, ok := token.Claims.(*user.TokenClaims); ok && token.Valid {
			c.Set("user", claims)
			c.Next()
		} else {
			http.Redirect(c.Writer, c.Request, "/auth/google/login", http.StatusTemporaryRedirect)
			return
		}
	}
}