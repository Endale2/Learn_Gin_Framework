package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/Learn_Gin_Framework/models"
)

// GET /books?author=X&minprice=Y&maxprice=Z
func GetAllBooks(c *gin.Context) {
	author := c.Query("author")
	minPriceStr := c.Query("minprice")
	maxPriceStr := c.Query("maxprice")

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

	var filtered []models.Book
	for _, book := range models.Books {
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

	if len(filtered) == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "No books match the filter"})
		return
	}

	c.JSON(http.StatusOK, filtered)
}

func GetBookByID(c *gin.Context) {
	strId := c.Param("id")
	Id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var foundBook *models.Book
	for _, book := range models.Books {
		if Id == book.ID {
			foundBook = &book
			break
		}
	}

	if foundBook == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, foundBook)
}

func CreateNewBook(c *gin.Context) {
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}
	if book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author is required"})
		return
	}
	if book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be above zero"})
		return
	}

	book.ID = len(models.Books) + 1
	models.Books = append(models.Books, book)

	c.JSON(http.StatusOK, gin.H{"msg": "Book successfully created", "book": book})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if id <= 0 || id > len(models.Books) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	models.Books = append(models.Books[:id-1], models.Books[id:]...)
	c.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted"})
}
