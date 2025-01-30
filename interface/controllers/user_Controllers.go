package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
    "github.com/apirestgo/infrastructure/database"
    "net/http"
    "strconv" // Necesitamos importar strconv
)

var db, _ = database.Connect()
var userRepository = &repository.UserRepository{DB: db}

func CreateUser(c *gin.Context) {
    var user domain.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := userRepository.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
    idStr := c.Param("id")  // Obtén el ID como string
    id, err := strconv.ParseUint(idStr, 10, 32)  // Convierte el string a uint64
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    user, err := userRepository.GetUserByID(uint(id))  // Convierte el ID a uint antes de usarlo
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    idStr := c.Param("id")  // Obtén el ID como string
    id, err := strconv.ParseUint(idStr, 10, 32)  // Convierte el string a uint64
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    
    var user domain.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.ID = uint(id)  // Asigna el ID convertido a uint

    if err := userRepository.UpdateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    idStr := c.Param("id")  // Obtén el ID como string
    id, err := strconv.ParseUint(idStr, 10, 32)  // Convierte el string a uint64
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    
    if err := userRepository.DeleteUser(uint(id));  // Convierte el ID a uint antes de usarlo
    err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
