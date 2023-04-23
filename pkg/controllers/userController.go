package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetUserEmail(c *gin.Context) string {
	user, _ := c.Get("user")
	return fmt.Sprint(user)
}

//  function for user registration

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Username string
		Type     string
	}

	if c.Bind(&body) != nil { // we parse json file with bind method and assign it to body fileds
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10) // we hash the password by bcrypt
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	user := models.User{Type: body.Type, Username: body.Username, Email: body.Email, Password: hashedPass} // create new user with body fields
	result := initializers.GetDB().Create(&user)                                                           // inserting user to database
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{}) // return status OK int json file
}

var Secret_Word = "HelloWorld"

func Get_Secret() string {
	return Secret_Word
}

//  for sign in user

func Login(c *gin.Context) {
	var body struct {
		Email    string // target gmail
		Password string // target password
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user := models.User{}
	initializers.GetDB().First(&user, "email = ?", body.Email) // getting first user with target gmail from users

	if user.ID == 0 { // if user doesn't exists
		c.JSON(http.StatusBadRequest, gin.H{ // return error
			"error": "Wrong gmail or password",
		})
		return
	}
	// else
	// comparing users password from table and hashed password from target password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong password. Try again",
		})
		return
	}
	// if passwords are the same then we create new claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": "jwt",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})
	// then new token
	tokenString, err := token.SignedString([]byte(Secret_Word))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a new token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authentication", tokenString, 3600*24*30, "", "", false, true) // set cookie on website
	c.JSON(http.StatusOK, gin.H{})
}

// for signing out

func SignOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true) // we need just set the cookie to expired one
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
