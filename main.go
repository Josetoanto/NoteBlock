package main

import (
	"database/sql"

	"github.com/apirestgo/application/services"
	"github.com/apirestgo/application/usecases"
	"github.com/apirestgo/infrastructure/handlers"
	"github.com/apirestgo/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Conexión a la base de datos
    db, err := sql.Open("mysql", "root:tono1234@tcp(localhost:3306)/notes_db")
    if err != nil {
        panic(err)
    }

    // Crear los repositorios y casos de uso
    noteRepository := persistence.NewNoteRepository(db) // Aquí pasamos la interfaz directamente
    createNoteUseCase := usecases.NewCreateNoteUseCase(noteRepository)
    getNotesUseCase := usecases.NewGetNotesUseCase(noteRepository)
    updateNoteUseCase := usecases.NewUpdateNoteUseCase(noteRepository)
    deleteNoteUseCase := usecases.NewDeleteNoteUseCase(noteRepository)

    userRepository := persistence.NewUserRepository(db)
    userService := services.NewUserService(userRepository)
    userHandler := handlers.NewUserHandler(userService)

    r.POST("/register", userHandler.RegisterUser)
    r.POST("/login", userHandler.LoginUser)

    // Crear el servicio
    noteService := services.NewNoteService(
        noteRepository, // Aquí pasamos la interfaz directamente
        createNoteUseCase,
        getNotesUseCase,
        updateNoteUseCase,
        deleteNoteUseCase,
    )

    // Crear el handler
    noteHandler := handlers.NewNoteHandler(noteService)

    // Definir las rutas
    r.GET("/notes", noteHandler.GetNotes)
    r.POST("/notes/create", noteHandler.CreateNote)
    r.PUT("/notes/update", noteHandler.UpdateNote)
    r.DELETE("/notes/delete", noteHandler.DeleteNote)

    // Iniciar el servidor
    r.Run(":8080")
}
