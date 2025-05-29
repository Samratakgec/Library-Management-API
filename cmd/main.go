package main

import (
	"github.com/Samratakgec/library-management/config"
	"github.com/Samratakgec/library-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize DB
    config.ConnectDB()

    // Create router
    r := gin.Default()

    // Register routes
    routes.BookRoutes(r)
    routes.TransactionLogsRoutes(r)

    // Start server
    r.Run(":8080")
}
