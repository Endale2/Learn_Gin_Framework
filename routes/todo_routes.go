package routes

import (
	"github.com/Endale2/Learn_Gin_Framework/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	
	api := r.Group("/api/v1/todos")
	{
		api.GET("/", controllers.GetTodos)           
		api.GET("/:id", controllers.GetTodoById)    
		api.POST("/", controllers.AddTodo)          
		api.PUT("/:id", controllers.UpdateTodo)     
		api.DELETE("/:id", controllers.DeleteTodo)  
	}
}
