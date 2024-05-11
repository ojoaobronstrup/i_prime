package repository

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/ojoaobronstrup/i_prime/entity"
)

type userRepository struct {
	db *sql.DB
}

type IRepository interface {
	FindUserByUsername(user entity.User) (bool, error)
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) FindUserByUsername(user entity.User) (bool, error) {
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
