package repository

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/ojoaobronstrup/i_prime/entity"
)

type validateUserRepository struct {
	db *sql.DB
}

func NewValidateUserRepository(db *sql.DB) *validateUserRepository {
	return &validateUserRepository{
		db: db,
	}
}

func (ur *validateUserRepository) GenerateToken(user entity.User) (string, error) {
	var isAdm bool
	err := ur.db.QueryRow("SELECT adm FROM users WHERE username = ?", user.Username).Scan(&isAdm)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("no user found: ", err)
			return "", err
		}
		slog.Error("error on execute the query: ", err)
		return "", err
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": user.Username,
			"sub": "validate_user",
			"adm": isAdm,
		},
	)

	if err := godotenv.Load("C:/Users/joaog/DEV/i_prime/.env"); err != nil {
		slog.Error("error on load the .env: ", err)
		return "", err
	}
	secretKey := os.Getenv("SECRET_KEY")
	key := []byte(secretKey)

	token, err := tokenClaims.SignedString(key)
	if err != nil {
		slog.Error("error on generate string token: ", err)
		return "", err
	}

	return token, nil
}

type IValidateUserRepository interface {
	GenerateToken(user entity.User) (string, error)
}
