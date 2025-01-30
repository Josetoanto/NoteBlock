package application

import (
    "github.com/apirestgo/domain"
    "github.com/apirestgo/infrastructure/repository"
)

type UpdateNoteUsecase struct {
    NoteRepository *repository.NoteRepository
}

func (uc *UpdateNoteUsecase) Execute(note *domain.Note) error {
    return uc.NoteRepository.UpdateNote(note)
}
