package usecase

import (
	"github.com/ojoaobronstrup/i_prime/entity"
	"github.com/ojoaobronstrup/i_prime/repository"
)

type UserUsecase struct {
	userRepository repository.IRepository
}

func NewUserUsecase(repo repository.IRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
	}
}

func (uc *UserUsecase) FindUserByUsername(user entity.User) (bool, error) {
	_, err := uc.userRepository.FindUserByUsername(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
