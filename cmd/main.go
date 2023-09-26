package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/domjeff/deploy-to-railway/internal/adapters/database"
	customHTTP "github.com/domjeff/deploy-to-railway/internal/adapters/http" // Alias the HTTP package
	"github.com/domjeff/deploy-to-railway/internal/application"

	"github.com/google/uuid"
)

func asal(interface{}) {}

func main() {
	// Initialize the PostgreSQL database
	db, err := database.InitPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close()

	a := uuid.NewString()

	asal(a)

	// Initialize application services
	userService := application.NewUserService(db)

	// Initialize HTTP handlers
	userHandler := customHTTP.UserHandler{UserService: userService} // Use the alias

	// Set up HTTP routes
	http.HandleFunc("/user", userHandler.CreateUserHandler)
	http.HandleFunc("/user-get", userHandler.GetUsersHandler)
	http.HandleFunc("/update-user", userHandler.UpdateUserHandler)

	// set up http routes for gorm
	http.HandleFunc("/gorm-student", userHandler.CreateGormStudentHandler)
	http.HandleFunc("/get-gorm-student", userHandler.GetGormStudentHandler)
	http.HandleFunc("/update-gorm-student", userHandler.UpdateGormStudentHandler)

	// Start the HTTP server
	serverPort := "8080" // Replace with your desired port
	fmt.Printf("Server listening on port %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
