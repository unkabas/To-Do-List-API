package main

import (
	"To-Do_List_Api/controllers"
	"To-Do_List_Api/initializers"
	"To-Do_List_Api/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Registration)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.LogOut)
	r.GET("/:username/profile", middlewares.CheckAuth, controllers.UserProfile)
	r.Run(":8080")
}
