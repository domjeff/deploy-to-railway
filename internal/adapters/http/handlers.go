// yourproject/internal/adapters/http/handlers.go

package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/domjeff/deploy-to-railway/internal/application"
	"github.com/domjeff/deploy-to-railway/internal/application/request"
	"github.com/domjeff/deploy-to-railway/internal/domain"
)

type UserHandler struct {
	UserService *application.UserService
}

// CreateUserHandler handles POST requests to create a new user.
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body into a User object
	var user request.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	u := domain.User{
		ID:       user.Id,
		UserName: user.Username,
		Email:    user.Email,
	}

	// Call the application service to create the user
	if err := h.UserService.CreateUser(&u); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func (h *UserHandler) CreateGormStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body into a User object
	var user request.CreateGormStudent
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	u := domain.Student{
		Email: user.Email,
	}

	// Call the application service to create the user
	if err := h.UserService.CreateGormStudent(&u); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func (h *UserHandler) GetGormStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid HTTP method", http.StatusNotAcceptable)
		return
	}

	student, err := h.UserService.GetGormStudent(id)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (h *UserHandler) UpdateGormStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid HTTP method", http.StatusNotAcceptable)
		return
	}

	// Parse the JSON request body into a User object
	var user request.UpdateGormStudent
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = h.UserService.UpdateGormStudent(id, user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	// Call the application service to create the user
	user, err := h.UserService.GetUsers()
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByIDHandler handles GET requests to retrieve a user by ID.
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body into a User object
	var user request.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	u := domain.User{
		ID:       user.Id,
		UserName: user.Username,
		Email:    user.Email,
	}

	// Call the application service to create the user
	if err := h.UserService.UpdateUser(&u); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}
