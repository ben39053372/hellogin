package controllers

import (
	"fmt"
	"hellogin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers get all user data
func GetAllUsers(c *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser create user
func CreateUser(c *gin.Context) {
	var user models.User
	fmt.Println("context: ", *c)
	c.BindJSON(&user)
	fmt.Println("user: ", user)
	err := models.CreateUsers(&user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
