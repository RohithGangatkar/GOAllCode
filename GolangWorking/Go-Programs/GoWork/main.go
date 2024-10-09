package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Year   int
}

var books = []Book{
	{
		Id: 1, Title: "book1", Author: "Author1", Year: 2019,
	},
	{
		Id: 1, Title: "book1", Author: "Author1", Year: 2019,
	},
}

func GetBooks(c *gin.Context) {

	c.JSON(200, books)

}

func AddBooks(c *gin.Context) {

	var book Book
	err := c.BindJSON(&book)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	books = append(books, book)

	c.JSON(http.StatusCreated, books)
}

func getBook(c *gin.Context) {

	id := c.Param("id")
	value, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	for _, eachBook := range books {
		if value == eachBook.Id {
			c.JSON(http.StatusOK, eachBook)
		}
	}

}

func main() {

	r := gin.Default()

	r.GET("v1/book", GetBooks)
	r.GET("v1/books/:name", getBook)
	r.POST("v1/book", AddBooks)
	r.Run(":8080")

}
