package main

import (
	"To-Do_List_Api/initializers"
	"To-Do_List_Api/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}
func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Task{})
}
