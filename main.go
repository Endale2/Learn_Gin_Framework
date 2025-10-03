package  main


import  "github.com/gin-gonic/gin"

var books []Book

type  Book  struct {
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
  books = append(books, book)
	c.JSON(200, gin.H{"message":"Book  Received", "book":book})
}

func  GetAllBooks(  c  *gin.Context){
   c.JSON(200, books)
}



func  main(){

	r:=gin.Default()

	r.POST("/create", CreateBooks)
	r.GET("/books", GetAllBooks)

	r.Run(":8080")
}