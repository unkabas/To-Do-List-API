package middlewares

import (
	"To-Do_List_Api/initializers"
	"To-Do_List_Api/models"
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	var user models.User
	username := c.Param("username")

	if err := initializers.DB.Where("username = ? AND is_auth = ?", username, true).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "you are not signin",
		})
		c.Abort()
		return
	}

	c.Set("currentUser", username)

	c.Next()
}
