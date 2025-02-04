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

    db, err := sql.Open("mysql", "root:tono1234@tcp(localhost:3306)/notes_db")
    if err != nil {
        panic(err)
    }

    noteRepository := persistence.NewNoteRepository(db) 
    createNoteUseCase := usecases.NewCreateNoteUseCase(noteRepository)
    getNotesUseCase := usecases.NewGetNotesUseCase(noteRepository)
    updateNoteUseCase := usecases.NewUpdateNoteUseCase(noteRepository)
    deleteNoteUseCase := usecases.NewDeleteNoteUseCase(noteRepository)

    userRepository := persistence.NewUserRepository(db)
    userService := services.NewUserService(userRepository)
    userHandler := handlers.NewUserHandler(userService)

    r.POST("/register", userHandler.RegisterUser)
    r.POST("/login", userHandler.LoginUser)

    noteService := services.NewNoteService(
        noteRepository, 
        createNoteUseCase,
        getNotesUseCase,
        updateNoteUseCase,
        deleteNoteUseCase,
    )

    noteHandler := handlers.NewNoteHandler(noteService)

    r.GET("/notes", noteHandler.GetNotes)
    r.POST("/notes/create", noteHandler.CreateNote)
    r.PUT("/notes/update", noteHandler.UpdateNote)
    r.DELETE("/notes/delete", noteHandler.DeleteNote)

    r.Run(":8080")
}
