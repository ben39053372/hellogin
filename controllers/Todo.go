package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTodos list all todo
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "list all todo",
	})
}

// CreateTodo create todo
func CreateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "create todo",
	})
}

// GetTodo get one todo item
func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "get one todo item",
	})
}

// UpdateTodo update todo item
func UpdateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "update todo item",
	})
}

// DeleteTodo delete todo item
func DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "delete todo item",
	})
}
