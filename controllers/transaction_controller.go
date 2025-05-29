package controllers

import (
	//"net/http"

	"github.com/Samratakgec/library-management/models"
	"github.com/Samratakgec/library-management/services"
	"github.com/gin-gonic/gin"
)

func GetTransactionLogByIsbnAndStdID(c *gin.Context) {
	isbn := c.Query("isbn")
	std_id := c.Query("std_id")

	if isbn == "" || std_id == "" {
		c.JSON(400, gin.H{"error": "isbn and std_id are required query parameters"})
		return
	}

	result, err := services.GetTransactionLogByIsbnAndStdID(isbn, std_id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, gin.H{"error": "no active allocation"})
			return
		} else {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
	}
	if result != nil {
		c.JSON(200, result)
	}
}
func AllocateBook(c *gin.Context) {
	var txn models.TransactionLogs
	if err := c.BindJSON(&txn); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}
	if book, err := services.GetBookByIsbn(txn.ISBN); book == nil || err != nil {
		if book == nil {
			c.JSON(404, gin.H{"error": "book with this isbn doesn't exist !!"})
			return
		} else if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
	}
	canAllot, err := services.CanBookBeAlloted(txn.StudentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	if !canAllot {
		c.JSON(409, gin.H{"error": "a book is already allocated to student"})
		return
	}

	if err := services.AllocateBook(txn); err != nil {

		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return

	}
	c.JSON(200, gin.H{"success": "book allocated successfully"})

}

func DeAllocateBook(c *gin.Context) {
	var txn models.TransactionLogs
	if err := c.BindJSON(&txn); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}
	result, err := services.GetTransactionLogByIsbnAndStdID(txn.ISBN, txn.StudentID)

	if result == nil {
		c.JSON(404, gin.H{"error": "no active allocation found"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := services.DeAllocateBook(txn); err != nil {

		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return

	}
	c.JSON(200, gin.H{"success": "book de-allocated successfully"})
}
