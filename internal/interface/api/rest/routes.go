package server

import (
	"blog-post/internal/domain/entity"
	"blog-post/internal/interface/api/rest/request"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := mux.NewRouter()
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/users", s.CreateUser).Methods("POST")

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Fail to decode request body", http.StatusBadRequest)
		return
	}

	user := entity.NewUser(userRequest.Name, userRequest.Email, userRequest.Password)
	log.Printf("New user created: %+v", user)
	// Call the Create method of SQLUserRepository to store the user in the database
	if err := s.service.CreateUser(user); err != nil {
		log.Printf("Failed to create user: %v", err) // Log the specific error
		http.Error(w, "Failed to create user, err:", http.StatusInternalServerError)
		return
	}
	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))

}
