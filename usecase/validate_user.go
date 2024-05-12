package usecase

import (
	"github.com/ojoaobronstrup/i_prime/entity"
	"github.com/ojoaobronstrup/i_prime/repository"
)

type ValidateUserUsecase struct {
	userRepository *repository.ValidateUserRepository
}

func NewValidateUserUsecase(repo *repository.ValidateUserRepository) *ValidateUserUsecase {
	return &ValidateUserUsecase{
		userRepository: repo,
	}
}

func (uc *ValidateUserUsecase) FindUserByUsername(user entity.User) (bool, error) {
	_, err := uc.userRepository.FindUserByUsername(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
