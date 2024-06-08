package rest

import (
	"blog-post/internal/domain/entity"
	"blog-post/internal/domain/repository"
	"blog-post/internal/service"
	"blog-post/internal/util/request"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func UserRouter(apiRouter *mux.Router, db *sqlx.DB) {

	// Create a subrouter for user routes
	userRouter := apiRouter.PathPrefix("/users").Subrouter()

	userRepository := repository.NewSQLUserRepository(db)
	userService := service.NewUserService(userRepository)
	// Define user routes with the appropriate paths
	userRouter.HandleFunc("", createUser(userService)).Methods("POST")
	userRouter.HandleFunc("/{userID}", getUserById(userService)).Methods("GET")
	userRouter.HandleFunc("", getAllUsers(userService)).Methods("GET")
}

func createUser(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest request.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			http.Error(w, "Fail to decode request body", http.StatusBadRequest)
			return
		}

		user := entity.NewUser(userRequest.Name, userRequest.Email, userRequest.Password)
		if err := userService.CreateUser(user); err != nil {
			log.Fatalf("Failed to create user: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully"))
	}
}

func getUserById(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.Parse(mux.Vars(r)["userID"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			log.Fatalf("Failed to get user by ID: %v", err)
			http.Error(w, fmt.Sprintf("Failed to get user by ID: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func getAllUsers(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		if err != nil {
			log.Fatalf("Failed to get all users: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
