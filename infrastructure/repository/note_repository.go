package repository

import (
    "github.com/apirestgo/domain"
    "github.com/jinzhu/gorm"
)

type NoteRepository struct {
    DB *gorm.DB
}

func (r *NoteRepository) CreateNote(note *domain.Note) error {
    return r.DB.Create(note).Error
}

func (r *NoteRepository) GetNoteByID(id int) (*domain.Note, error) {
    var note domain.Note
    if err := r.DB.First(&note, id).Error; err != nil {
        return nil, err
    }
    return &note, nil
}

func (r *NoteRepository) UpdateNote(note *domain.Note) error {
    return r.DB.Save(note).Error
}

func (r *NoteRepository) DeleteNote(id int) error {
    return r.DB.Delete(&domain.Note{}, id).Error
}
