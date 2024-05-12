package repository

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDatabase() (*sql.DB, error) {
	if err := godotenv.Load("C:/Users/joaog/DEV/i_prime/.env"); err != nil {
		slog.Error("error on load the env file: ", err)
		return nil, err
	}

	dbAccess := os.Getenv("DATABASE_ACCESS")

	db, err := sql.Open("mysql", dbAccess)
	if err != nil {
		slog.Error("database connection error: ", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		slog.Error("database connection error: ", err)
		return nil, err
	}

	return db, nil
}
