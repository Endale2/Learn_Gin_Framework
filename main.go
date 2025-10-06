package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


var  books []Book
type  Book  struct{
	ID  int  `json:"id"`
	Title string  `json:"title"`
	Author string `json:"author"`
	Price   int  `json:"price"`
}

func GetAllBooks(c *gin.Context) {
    author := c.Query("author")
    minPriceStr := c.Query("minprice")
    maxPriceStr := c.Query("maxprice")

    // Convert prices to integers (if provided)
    var minPrice, maxPrice int
    var err error

    if minPriceStr != "" {
        minPrice, err = strconv.Atoi(minPriceStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "minprice must be a number"})
            return
        }
    }

    if maxPriceStr != "" {
        maxPrice, err = strconv.Atoi(maxPriceStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "maxprice must be a number"})
            return
        }
    }

    // Filter books
    var filtered []Book
    for _, book := range books {
        if author != "" && book.Author != author {
            continue
        }
        if minPriceStr != "" && book.Price < minPrice {
            continue
        }
        if maxPriceStr != "" && book.Price > maxPrice {
            continue
        }
        filtered = append(filtered, book)
    }

    // Response
    if len(filtered) == 0 {
        c.JSON(http.StatusOK, gin.H{"msg": "No books match the filter"})
        return
    }

    c.JSON(http.StatusOK, filtered)
}

func  GetBookByID(c *gin.Context){
	strId:=c.Param("id")
	Id, err:=strconv.Atoi(strId)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
		return
	}

	var  foundBook *Book

	for _, book:= range books{
		if Id==book.ID{
			foundBook=&book
			break
		}
	}

	if  foundBook==nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"The book is Not Found!"})
		return
	}

	c.JSON(http.StatusOK, foundBook)
}


func  CreateNewBook(c *gin.Context){
	var  book Book

	err:=c.BindJSON(&book)
	if err!=nil{
      c.JSON(http.StatusBadRequest, gin.H{"error":"Error  Binding JSON"})
	  return
	}

	if book.Title=="" {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Title  cannot  be  empty"})
		return
	}

	if book.Author ==""{
      c.JSON(http.StatusBadRequest, gin.H{"error":"Author  is  required"})
	  return
	}

	if  book.Price <=0{
      c.JSON(http.StatusBadRequest, gin.H{"error":"Price  should  be  above  zero"})
	}

	book.ID=len(books)+1
	books =append(books,book)

	c.JSON(http.StatusOK, gin.H{"msg":"the  book is  successfully  created","book":book})
}


func  DeleteBook(c  *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid  ID"})
		return
	}
  if  id<=0 ||id>len(books){
	c.JSON(http.StatusOK, gin.H{"error":"Book  Not  Found"})
	return
  }
   books =append(books[:id-1], books[id+ 1:]...)

   c.JSON(http.StatusOK, gin.H{"msg":"successfully  Deleted"})
}

//middleware

func  BookMiddleware(c  *gin.Context){
	if len(books)==0{
		c.JSON(404, gin.H{"msg":"no books  yet"})
		c.Abort()
	}
	c.Next()
}

func  main(){

	r:=gin.New()
	r.Use(gin.Logger())


	r.GET("/books",BookMiddleware, GetAllBooks)
	r.GET("/books/:id", GetBookByID)
	r.POST("/books", CreateNewBook)
	r.DELETE("/books/:id", DeleteBook)


	r.Run(":8080")

}