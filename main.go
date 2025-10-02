package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    
    r := gin.New()
	r.Use(gin.Logger())

    
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
	

    // start server on port 8080
    r.Run(":8080")
}
