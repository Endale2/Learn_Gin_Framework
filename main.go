package main

import (


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
	r.GET("users/:id", func (ctx *gin.Context){
		id:=ctx.Param("id")

		ctx.JSON(200, gin.H{"user_id":id})

	})

	r.GET("/search", func(c *gin.Context){
		name:=c.Query("name")
		age:=c.DefaultQuery("age", "0")
       resp:=gin.H{"name":name, "age":age}
		c.JSON(200, resp)
	})

    // start server on port 8080
    r.Run(":8080")
}
