package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Endale2/Learn_Gin_Framework/controllers"
	"github.com/Endale2/Learn_Gin_Framework/middleware"
)

func BookRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	books := v1.Group("/books")

	books.GET("/", middleware.BookMiddleware, controllers.GetAllBooks)
	books.GET("/:id", controllers.GetBookByID)
	books.POST("/", controllers.CreateNewBook)
	books.DELETE("/:id", controllers.DeleteBook)
}
