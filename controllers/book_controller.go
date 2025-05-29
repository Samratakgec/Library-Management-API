package controllers

import (
	//"net/http"

	"github.com/Samratakgec/library-management/models"
	"github.com/Samratakgec/library-management/services"
	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "In-valid book payload"})
	}

	// Check if a book with the same ISBN already exists
	b1, err := services.GetBookByIsbn(book.ISBN)
	if err == nil && b1 != nil {
		// Book exists
		c.JSON(409, gin.H{"error": "Book with this ISBN already exists!!"})
		return
	}
	if err != nil && err.Error() != "book not found" {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// adding a book
	if err := services.AddBook(book); err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(200, gin.H{"success": "book added successfully"})
}

func GetBookByIsbn(c *gin.Context) {
	isbn := c.Param("isbn")
	if isbn == "" {
		c.JSON(400, gin.H{"error": "please provide isbn"})
	}
	book, err := services.GetBookByIsbn(isbn)
	if err.Error() == "book not found" {
		c.JSON(404, gin.H{"error": "book with the provided isbn does not exist"})
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(200, gin.H{"book": book})
}

func DeleteBookByIsbn(c *gin.Context) {
	isbn := c.Param("isbn")
	if isbn == "" {
		c.JSON(400, gin.H{"error": "please provide isbn"})
	}
	err := services.DeleteBookByIsbn(isbn)
	if err == nil {
		c.JSON(200, gin.H{"success": "book deleted successfully"})
	} else if err.Error() == "book not found" {
		c.JSON(404, gin.H{"error": "book with the provided isbn does not exist"})
	}
}
