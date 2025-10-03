package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books []Book

type  Book  struct {
	Id  int  `json:"id"`
	Title string  `json:"title"`
	Author string `json:"author"`
	Publisher  Publisher `json:"publisher"`
}

type  Publisher struct{
	Name  string  `json:"name"`
	City  string  `json:"city"`
}

func  CreateBooks(c  *gin.Context){
	var  book  Book
	err:=c.BindJSON(&book);

	if  err!=nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	book.Id  = len(books) +1
    books = append(books, book)
  
	c.JSON(200, gin.H{"message":"Book  Received", "book":book})
}

func  GetAllBooks(  c  *gin.Context){
   c.JSON(200, books)
}


func  GetBook(  c *gin.Context){
	idstr:=c.Param("id")
id,err:=strconv.Atoi(idstr)
if  err!=nil{
	c.JSON(http.StatusBadRequest, map[string]any{"error":"Invalid ID"})
}
	var  foundBook  *Book 

	for _,  book:=range books{
		if  book.Id==id{
			foundBook = &book
			break
		}
	}

	if foundBook==nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"Book  Not  Found"})
		return
	}

c.JSON(http.StatusOK, foundBook)
}


func  main(){

	r:=gin.Default()

	r.POST("/books", CreateBooks)
	r.GET("/books", GetAllBooks)
	r.GET("/books/:id", GetBook)

	r.Run(":8080")
}