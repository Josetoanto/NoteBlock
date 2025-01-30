package application

import (
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
)

type CreateNoteUsecase struct {
    NoteRepository *repository.NoteRepository
}

func (uc *CreateNoteUsecase) Execute(note *domain.Note) error {
    return uc.NoteRepository.CreateNote(note)
}
