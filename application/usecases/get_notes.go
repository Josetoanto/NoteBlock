package usecases

import (
    "github.com/apirestgo/domain/entities"
    "github.com/apirestgo/infrastructure/persistence"
)

type GetNotesUseCase struct {
    repository persistence.NoteRepository
}

func NewGetNotesUseCase(repository persistence.NoteRepository) *GetNotesUseCase {
    return &GetNotesUseCase{repository: repository}
}

func (uc *GetNotesUseCase) Execute(userID int) ([]entities.Note, error) {
    return uc.repository.GetByUserID(userID) // Cambiado de GetNotesByUserID a GetByUserID
}
