package middleware

import (
	"Test2/initializers"
	"Test2/pkg/controllers"
	"Test2/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// checking users for authentication

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authentication") // get Authentication cookie
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { // get token by parsing the cookie
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controllers.Get_Secret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) { // statement for claims expiration
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.GetDB().First(&user, claims["sub"]) // getting user by claims
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user.Email)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
