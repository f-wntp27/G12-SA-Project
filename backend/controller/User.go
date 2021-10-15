package controller

import (
	"net/http"
	"github.com/fonthap/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []entity.User
	if err := entity.DB().Raw("SELECT * FROM users").Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
