package rest

import (
	"blog-post/internal/domain/entity"
	"blog-post/internal/interface/api/rest/request"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) UserRoutes(apiRouter *mux.Router) {
	// Create a subrouter for user routes
	userRouter := apiRouter.PathPrefix("/users").Subrouter()

	// Define user routes with the appropriate paths
	userRouter.HandleFunc("", s.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", s.GetUserById).Methods("GET")
	userRouter.HandleFunc("", s.GetAllUsers).Methods("GET")
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Fail to decode request body", http.StatusBadRequest)
		return
	}

	user := entity.NewUser(userRequest.Name, userRequest.Email, userRequest.Password)
	// Call the Create method of SQLUserRepository to store the user in the database
	if err := s.service.CreateUser(user); err != nil {
		log.Fatalf("Failed to create user : %v", err)
		http.Error(w, "Failed to create user, err:", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))

}

func (s *Server) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := s.service.GetUserByID(userID)
	if err != nil {
		log.Fatalf("Failed to get user by ID: %v", err)
		http.Error(w, fmt.Sprintf("Failed to get user by ID: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.GetAllUsers()
	if err != nil {
		log.Fatalf("Failed to get all users: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
