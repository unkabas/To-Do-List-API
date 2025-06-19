package controllers

import (
	"To-Do_List_Api/initializers"
	"To-Do_List_Api/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(400, gin.H{
			"message": "Not a JSON",
		})
		return
	}
	if authInput.Username == "" || authInput.Password == "" {
		c.JSON(400, gin.H{
			"message": "No username or password added",
		})
		return
	}
	var userFound models.User
	if err := initializers.DB.Where("username=? AND password=?", authInput.Username, authInput.Password).First(&userFound).Error; err != nil {
		c.JSON(400, gin.H{
			"message":  "fuck",
			"username": authInput.Username,
		})
	} else {
		initializers.DB.Model(&userFound).Update("IsAuth", "True")
		c.JSON(200, gin.H{
			"message":  "success",
			"username": authInput.Username,
		})
	}

}
