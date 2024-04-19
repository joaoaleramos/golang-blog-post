package rest

import (
	"blog-post/internal/domain/repository"
	"blog-post/internal/service"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface{}

type Server struct {
	port        int
	userService *service.UserService
	postService *service.PostService
}

func NewServer(db *sqlx.DB) *http.Server {
	sqlUserRepository := repository.NewSQLUserRepository(db)
	sqlPostRepository := repository.NewSQLPostRepository(db)
	userService := service.NewUserService(sqlUserRepository)
	postService := service.NewPostService(sqlPostRepository)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:        port,
		userService: userService,
		postService: postService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
