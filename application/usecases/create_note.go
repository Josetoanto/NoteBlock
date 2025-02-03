package usecases

import (

	"github.com/apirestgo/domain/entities"
	"github.com/apirestgo/infrastructure/persistence"
)

type CreateNoteUseCase struct {
	repository persistence.NoteRepository
}

// NewCreateNoteUseCase crea una nueva instancia de CreateNoteUseCase.
func NewCreateNoteUseCase(repository persistence.NoteRepository) *CreateNoteUseCase {
	return &CreateNoteUseCase{repository: repository}
}

// Execute ejecuta el caso de uso para crear una nueva nota.
func (uc *CreateNoteUseCase) Execute(userID int, title, description string) error {
	

	note := &entities.Note{
		Title:       title,
		Description: description,
		UserID:      userID,    // Asociar el user_id con la nota
	}

	return uc.repository.Create(note)
}
