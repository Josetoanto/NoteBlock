package usecases

import (
    "github.com/apirestgo/domain/entities"
    "github.com/apirestgo/infrastructure/persistence"
)

type UpdateNoteUseCase struct {
    repository persistence.NoteRepository
}

func NewUpdateNoteUseCase(repository persistence.NoteRepository) *UpdateNoteUseCase {
    return &UpdateNoteUseCase{repository: repository}
}

func (uc *UpdateNoteUseCase) Execute(note *entities.Note) error {
    return uc.repository.Update(note)
}
