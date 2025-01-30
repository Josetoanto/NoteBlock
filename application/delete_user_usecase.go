package application

import "github.com/apirestgo/infrastructure/repository"

type DeleteUserUsecase struct {
    UserRepository *repository.UserRepository
}

func (uc *DeleteUserUsecase) Execute(id uint) error {
    return uc.UserRepository.DeleteUser(id)
}
