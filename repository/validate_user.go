package repository

import (
	"database/sql"
	"log"
	"log/slog"

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

func (ur *validateUserRepository) FindUserByUsername(user entity.User) (bool, error) {
	var foundUsername string
	err := ur.db.QueryRow("SELECT username FROM users WHERE username = ?", user.Username).Scan(&foundUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("no user found: ", err)
			return false, err
		}
		slog.Error("error on execute the query: ", err)
		return false, err
	}

	log.Println(user.Username)

	return true, nil
}

type IValidateUserRepository interface {
	FindUserByUsername(user entity.User) (bool, error)
}
