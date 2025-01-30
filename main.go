package main

import (
    "github.com/gin-gonic/gin"
    "github.com/apirestgo/interface/controllers"
)

func main() {
    r := gin.Default()

    // Inicializar las rutas
    controllers.InitializeRoutes(r)

    // Iniciar el servidor en el puerto 8080
    r.Run(":8080")
}
