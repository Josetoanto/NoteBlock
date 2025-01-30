package controllers

import (
	"net/http"
	"strconv"

	"github.com/apirestgo/application"
	"github.com/apirestgo/domain"
	"github.com/apirestgo/infrastructure/database" // Importa el paquete de database
	"github.com/apirestgo/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

var noteRepository = &repository.NoteRepository{DB: database.DB}
var createNoteUsecase = &application.CreateNoteUsecase{NoteRepository: noteRepository}
var getNoteUsecase = &application.GetNoteUsecase{NoteRepository: noteRepository}
var updateNoteUsecase = &application.UpdateNoteUsecase{NoteRepository: noteRepository}
var deleteNoteUsecase = &application.DeleteNoteUsecase{NoteRepository: noteRepository}

func CreateNote(c *gin.Context) {
    var note domain.Note
    if err := c.BindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := createNoteUsecase.Execute(&note); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, note)
}

func GetNote(c *gin.Context) {
	idStr := c.Param("id")  // Obtén el ID como string
    id, err := strconv.Atoi(idStr) // Convierte el string a int
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    note, err := getNoteUsecase.Execute(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }
    c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
    idStr := c.Param("id")  // Obtén el ID como string
    id, err := strconv.Atoi(idStr)  // Convierte el string a int
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var note domain.Note
    if err := c.BindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Convierte el int a el ID de la nota
    note.ID = id

    if err := updateNoteUsecase.Execute(&note); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	idStr := c.Param("id")  // Obtén el ID como string
    id, err := strconv.Atoi(idStr) // Convierte el string a int
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := deleteNoteUsecase.Execute(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
