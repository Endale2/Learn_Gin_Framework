package main

import (
	"github.com/Endale2/Learn_Gin_Framework/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	routes.BookRoutes(r)

	r.Run(":8080")
}
