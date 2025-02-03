package persistence

import (
    "database/sql"
    "github.com/apirestgo/domain/entities"
    "golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    query := "INSERT INTO users (username, password) VALUES (?, ?)"
    _, err = r.DB.Exec(query, user.Username, hashedPassword)
    return err
}

func (r *UserRepository) FindUserByUsername(username string) (*entities.User, error) {
    var user entities.User
    query := "SELECT id, username, password FROM users WHERE username = ?"
    err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}


