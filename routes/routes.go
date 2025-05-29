package routes

import (
	"github.com/Samratakgec/library-management/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes (router *gin.Engine){
	// group-all book related functions
	book := router.Group("/book")
	{
		book.POST("/new",controllers.AddBook)
		book.GET("/:isbn",controllers.GetBookByIsbn)
		book.DELETE("/:isbn",controllers.DeleteBookByIsbn)
	}
}

func TransactionLogsRoutes(router *gin.Engine)  {
	transaction := router.Group("/transaction")
	{
		transaction.POST("/allocate",controllers.AllocateBook)
		transaction.GET("/get-log", controllers.GetTransactionLogByIsbnAndStdID)
		transaction.PUT("/de-allocate",controllers.DeAllocateBook)
	}
}