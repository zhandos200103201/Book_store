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

// function for defining role of users

func UserIsAdmin(c *gin.Context) {
	tokenString, _ := c.Cookie("Authentication")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controllers.Get_Secret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.GetDB().First(&user, claims["sub"])
		if user.Type != "Admin" { // checking for admins role
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Only Admins have the permission",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func UserIsClientOrAdmin(c *gin.Context) {
	tokenString, _ := c.Cookie("Authentication")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controllers.Get_Secret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.GetDB().First(&user, claims["sub"])
		if user.Type != "Client" && user.Type != "Admin" { // checking for admins or client role
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Only Admins and Clients have the permission",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func UserIsClient(c *gin.Context) {
	tokenString, _ := c.Cookie("Authentication")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controllers.Get_Secret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.GetDB().First(&user, claims["sub"])
		if user.Type != "Client" { // checking for clients role
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Only Admins and Clients have the permission",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func UserIsSellerOrAdmin(c *gin.Context) {
	tokenString, _ := c.Cookie("Authentication")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controllers.Get_Secret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.GetDB().First(&user, claims["sub"])
		if user.Type != "Seller" && user.Type != "Admin" { // checking for admins or sellers role
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Only Admins and Sellers have the permission",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func UserIsSeller(c *gin.Context) {
	tokenString, _ := c.Cookie("Authentication")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controllers.Get_Secret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.GetDB().First(&user, claims["sub"])
		if user.Type != "Seller" { // checking for sellers role
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Only Admins and Sellers have the permission",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
