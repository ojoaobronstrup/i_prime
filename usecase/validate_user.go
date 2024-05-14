package usecase

import (
	"github.com/ojoaobronstrup/i_prime/entity"
	"github.com/ojoaobronstrup/i_prime/repository"
)

type ValidateUserUsecase struct {
	userRepository repository.IValidateUserRepository
}

func NewValidateUserUsecase(repo repository.IValidateUserRepository) *ValidateUserUsecase {
	return &ValidateUserUsecase{
		userRepository: repo,
	}
}

func (uc *ValidateUserUsecase) GenerateToken(user entity.User) (string, error) {
	_, err := uc.userRepository.GenerateToken(user)
	if err != nil {
		return "user not found", err
	}
	return user.Username, nil
}
