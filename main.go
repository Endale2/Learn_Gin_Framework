package main

import (
	"github.com/Endale2/Learn_Gin_Framework/database"
	"github.com/Endale2/Learn_Gin_Framework/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()
	routes.TodoRoutes(r)

	r.Run(":8080")
}
