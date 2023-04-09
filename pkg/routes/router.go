package routes

import (
	"Test2/middleware"
	"Test2/pkg/controllers"
	gin "github.com/gin-gonic/gin"
)

var Router = func(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", middleware.RequireAuth, middleware.UserIsSeller, controllers.CreateBook)
	router.POST("/give_a_rating", middleware.RequireAuth, middleware.UserIsClient, controllers.GiveRating)
	router.POST("/comment", middleware.RequireAuth, middleware.UserIsClient, controllers.CreateComment)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", middleware.RequireAuth, middleware.UserIsSeller, controllers.EditBook)
	router.DELETE("/books/:id", middleware.RequireAuth, middleware.UserIsSeller, controllers.DeleteBook)
	router.GET("/filter_by_prices", controllers.PriceFiltering)
	router.GET("/filter_by_rating", controllers.RatingFiltering)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", middleware.RequireAuth, controllers.SignOut)
	err := router.Run()
	if err != nil {
		return
	}
}
