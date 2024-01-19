package booklib

import (
	"booklib/internal/database"
	"booklib/pkg/router"
)

func main() {
	// Initialize the database connection
	database.InitDB()

	// Create a new Gin router
	r := router.NewRouter()

	// Start the server
	r.Run(":8080")
}
