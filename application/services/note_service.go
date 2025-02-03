package services

import (
	"github.com/apirestgo/application/usecases"
	"github.com/apirestgo/domain/entities"
	"github.com/apirestgo/infrastructure/persistence"
)

type NoteService struct {
    repo              persistence.NoteRepository
    createNoteUseCase *usecases.CreateNoteUseCase
    getNotesUseCase   *usecases.GetNotesUseCase
    updateNoteUseCase *usecases.UpdateNoteUseCase
    deleteNoteUseCase *usecases.DeleteNoteUseCase
}

func NewNoteService(
    repo persistence.NoteRepository, // Cambiado de *persistence.NoteRepository a persistence.NoteRepository
    createNoteUseCase *usecases.CreateNoteUseCase,
    getNotesUseCase *usecases.GetNotesUseCase,
    updateNoteUseCase *usecases.UpdateNoteUseCase,
    deleteNoteUseCase *usecases.DeleteNoteUseCase,
) *NoteService {
    return &NoteService{
        repo:              repo,
        createNoteUseCase: createNoteUseCase,
        getNotesUseCase:   getNotesUseCase,
        updateNoteUseCase: updateNoteUseCase,
        deleteNoteUseCase: deleteNoteUseCase,
    }
}

func (s *NoteService) CreateNote(userID int, title string, description string) error {
	// Crear la nueva nota, asegurándose de asociarla con el userID
	note := entities.Note{
		Title:       title,
		Description: description,
		UserID:      userID, // Asociar la nota con el usuario
	}

	// Llamar al repositorio para guardar la nota
	return s.repo.Create(&note)
}

func (s *NoteService) GetNotes(userID int) ([]entities.Note, error) {
	// Obtener las notas del usuario desde el repositorio
	return s.repo.GetByUserID(userID)
}

func (s *NoteService) UpdateNote(note *entities.Note) error {
    return s.updateNoteUseCase.Execute(note)
}

func (s *NoteService) DeleteNote(id int) error {
    return s.deleteNoteUseCase.Execute(id)
}
