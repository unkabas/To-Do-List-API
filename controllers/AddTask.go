package controllers

import (
	"To-Do_List_Api/initializers"
	"To-Do_List_Api/models"
	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	username := c.Param("username")

	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "No such user",
		})
		return
	}

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{
			"message": "Not a JSON",
		})
		return
	}

	if newTask.Title == "" || newTask.Content == "" {
		c.JSON(400, gin.H{
			"message": "No Title or Content added",
		})
		return
	}

	newTask.UserID = username

	if err := initializers.DB.Create(&newTask).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Task not created",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
