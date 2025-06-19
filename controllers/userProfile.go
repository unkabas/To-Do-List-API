package controllers

import (
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})

}
