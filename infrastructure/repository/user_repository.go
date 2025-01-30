package repository

import (
    "github.com/apirestgo/domain"
    "github.com/jinzhu/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
    var user domain.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
    return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
    return r.DB.Delete(&domain.User{}, id).Error
}
