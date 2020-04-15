package routers

import (
	"hellogin/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter main router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		todo := v1.Group("/todo")
		{
			todo.GET("/", controllers.GetTodos)
			todo.POST("/", controllers.CreateTodo)
			todo.GET("/:id", controllers.GetTodo)
			todo.PUT("/:id", controllers.UpdateTodo)
			todo.DELETE("/:id", controllers.DeleteTodo)
		}
		user := v1.Group("/user")
		{
			user.GET("/", controllers.GetAllUsers)
			user.POST("/", controllers.CreateUser)
		}
	}

	return r
}
