package application

import (
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
)

type GetUserUsecase struct {
    UserRepository *repository.UserRepository
}

func (uc *GetUserUsecase) Execute(id uint) (*domain.User, error) {
    return uc.UserRepository.GetUserByID(id)
}
