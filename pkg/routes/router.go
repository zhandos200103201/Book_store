package routes

import (
	"Test2/middleware"
	"Test2/pkg/controllers"
	gin "github.com/gin-gonic/gin"
)

var Router = func(router *gin.Engine) {
	//route.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	//route.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	//route.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	//route.HandleFunc("/books/{id}", controllers.EditBook).Methods("PUT")
	//route.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/signout", controllers.SignOut)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	err := router.Run()
	if err != nil {
		return
	}
}
