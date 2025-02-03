package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/apirestgo/application/services"
    "net/http"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{userService: service}
}

// Registrar usuario
func (h *UserHandler) RegisterUser(c *gin.Context) {
    var user struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    
    if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}	

    err := h.userService.CreateUser(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Iniciar sesión y obtener token
func (h *UserHandler) LoginUser(c *gin.Context) {
    var user struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := h.userService.Authenticate(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
