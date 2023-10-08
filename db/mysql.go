package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DatabaseConnection struct {
	DB   *sql.DB
	Err  error
}
func Connect() DatabaseConnection {

	godotenv.Load()
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, databaseName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return DatabaseConnection{nil, err}
	}

	err = db.Ping()
	if err != nil {
		return DatabaseConnection{nil, err}
	}

	return DatabaseConnection{db, err}
}