package controller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/ojoaobronstrup/i_prime/entity"
	"github.com/ojoaobronstrup/i_prime/repository"
	"github.com/ojoaobronstrup/i_prime/usecase"
)

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	db, err := repository.InitDatabase()
	if err != nil {
		slog.Error("database connection error: ", err)
		return
	}

	userRepo := repository.NewValidateUserRepository(db)
	userUsecase := usecase.NewValidateUserUsecase(userRepo)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error on read the request payload: ", err)
		return
	}

	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		slog.Error("error on convert the payload into an user: ", err)
		return
	}

	_, err = userUsecase.GenerateToken(user)
	if err != nil {
		slog.Error("user not found: ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
