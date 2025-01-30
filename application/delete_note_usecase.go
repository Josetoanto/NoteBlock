package application

import "github.com/apirestgo/infrastructure/repository"

type DeleteNoteUsecase struct {
    NoteRepository *repository.NoteRepository
}

func (uc *DeleteNoteUsecase) Execute(id int) error {
    return uc.NoteRepository.DeleteNote(id)
}
