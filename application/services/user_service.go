package services

import (
    "github.com/apirestgo/domain/entities"
    "github.com/apirestgo/infrastructure/persistence"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
    "errors"
)

type UserService struct {
    userRepository *persistence.UserRepository
}

func NewUserService(userRepo *persistence.UserRepository) *UserService {
    return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(username, password string) error {
    user := &entities.User{
        Username: username,
        Password: password,
    }
    return s.userRepository.CreateUser(user)
}

func (s *UserService) Authenticate(username, password string) (string, error) {
    user, err := s.userRepository.FindUserByUsername(username)
    if err != nil {
        return "", err
    }
    if user == nil {
        return "", errors.New("user not found")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid password")
    }

    token, err := generateJWT(user.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}

func generateJWT(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secretKey := []byte("mysecretkey") // Usar un valor seguro para la clave secreta
    signedToken, err := token.SignedString(secretKey)
    return signedToken, err
}
