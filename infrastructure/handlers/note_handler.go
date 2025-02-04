package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/apirestgo/application/services"
	"github.com/apirestgo/domain/entities"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	service *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{service: service}
}

// GetNotes ahora usa *gin.Context
func (h *NoteHandler) GetNotes(c *gin.Context) {
	// Obtener el token del header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("mysecretkey"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID := int(claims["user_id"].(float64)) // El user_id es un número en el JWT

	notes, err := h.service.GetNotes(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var note entities.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("mysecretkey"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID := int(claims["user_id"].(float64)) 

	err = h.service.CreateNote(userID, note.Title, note.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Note created"})
}


func (h *NoteHandler) UpdateNote(c *gin.Context) {
	var note entities.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	err := h.service.UpdateNote(&note) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note updated"})
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	idStr := c.DefaultQuery("id", "") 
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	id, err := strconv.Atoi(idStr) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.service.DeleteNote(id) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Note deleted"})
}
