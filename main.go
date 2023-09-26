package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	r := http.NewServeMux()

	// Set up HTTP routes
	r.Handle("/ping", http.HandlerFunc(userHandler.Ping))
	r.Handle("/user", loggingMiddleware(http.HandlerFunc(userHandler.CreateUserHandler)))
	r.Handle("/user-get", loggingMiddleware(http.HandlerFunc(userHandler.GetUsersHandler)))
	r.Handle("/update-user", loggingMiddleware(http.HandlerFunc(userHandler.UpdateUserHandler)))
	r.Handle("/gorm-student", loggingMiddleware(http.HandlerFunc(userHandler.CreateGormStudentHandler)))
	r.Handle("/get-gorm-student", loggingMiddleware(http.HandlerFunc(userHandler.GetGormStudentHandler)))
	r.Handle("/update-gorm-student", loggingMiddleware(http.HandlerFunc(userHandler.UpdateGormStudentHandler)))

	// Start the HTTP server
	serverPort := "8080" // Replace with your desired port
	fmt.Printf("Server listening on port %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Endpoint: %s\n", r.URL.Path)

		next.ServeHTTP(w, r)

		func(w http.ResponseWriter, r *http.Request) {
			status, _ := strconv.Atoi(w.Header().Get("Status-Code"))
			if status < 200 || status >= 300 {
				fmt.Printf("Request to %s resulted in a non-successful status code: %d\n", r.URL.Path, status)
			}
		}(w, r)

	})
}
