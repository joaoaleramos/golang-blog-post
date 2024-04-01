package postgres

import (
	"blog-post/internal/utils"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	dbInstance *sqlx.DB
	once       sync.Once
	settings   utils.DBConfig = *utils.LoadDBConfig()
)

// GetDBInstance returns a singleton instance of the database connection
func GetDBInstance() (*sqlx.DB, error) {
	once.Do(func() {
		// Initialize the database connection
		db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=%v", settings.DbUser, settings.DbPassword, settings.DbName, settings.DbPort, settings.DbSSLMode))
		if err != nil {
			panic(err) // or handle the error gracefully
		}
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)

	})

	return dbInstance, nil
}
