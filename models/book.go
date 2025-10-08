package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

var Books []Book
