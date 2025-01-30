package controllers

import (
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
    r.POST("/products", CreateNote)
    r.GET("/products/:id", GetNote)
    r.PUT("/products/:id", UpdateNote)
    r.DELETE("/products/:id", DeleteNote)

    // Rutas para usuarios
    r.POST("/users", CreateUser)
    r.GET("/users/:id", GetUser)
    r.PUT("/users/:id", UpdateUser)
    r.DELETE("/users/:id", DeleteUser)
}
