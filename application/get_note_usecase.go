package application

import (
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
)

type GetNoteUsecase struct {
    NoteRepository *repository.NoteRepository
}

func (uc *GetNoteUsecase) Execute(id int) (*domain.Note, error) {
    return uc.NoteRepository.GetNoteByID(id)
}
