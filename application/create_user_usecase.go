package application

import (
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
)

type CreateUserUsecase struct {
    UserRepository *repository.UserRepository
}

func (uc *CreateUserUsecase) Execute(user *domain.User) error {
    return uc.UserRepository.CreateUser(user)
}
