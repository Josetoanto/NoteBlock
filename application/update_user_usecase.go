package application

import (
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
)

type UpdateUserUsecase struct {
    UserRepository *repository.UserRepository
}

func (uc *UpdateUserUsecase) Execute(user *domain.User) error {
    return uc.UserRepository.UpdateUser(user)
}
