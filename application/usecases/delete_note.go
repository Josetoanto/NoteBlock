package usecases

import "github.com/apirestgo/infrastructure/persistence"

type DeleteNoteUseCase struct {
    repository persistence.NoteRepository
}

func NewDeleteNoteUseCase(repository persistence.NoteRepository) *DeleteNoteUseCase {
    return &DeleteNoteUseCase{repository: repository}
}

func (uc *DeleteNoteUseCase) Execute(id int) error {
    return uc.repository.Delete(id)
}
