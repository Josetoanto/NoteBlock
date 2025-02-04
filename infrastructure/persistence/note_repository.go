package persistence

import (
	"database/sql"
	"log"

	"github.com/apirestgo/domain/entities"
	_ "github.com/go-sql-driver/mysql"
)

type NoteRepository interface {
	Create(note *entities.Note) error
	GetByUserID(userID int) ([]entities.Note, error)
	GetByID(id int) (*entities.Note, error)
	Update(note *entities.Note) error
	Delete(id int) error
}

type noteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) Create(note *entities.Note) error {
	query := "INSERT INTO notes (title, description, user_id) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, note.Title, note.Description, note.UserID)
	return err
}

func (r *noteRepository) GetByUserID(userID int) ([]entities.Note, error) {
	rows, err := r.db.Query("SELECT id, title, description, user_id FROM notes WHERE user_id = ?", userID)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		return nil, err
	}
	defer rows.Close()

	var notes []entities.Note
	for rows.Next() {
		var note entities.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Description, &note.UserID); err != nil {
			log.Println("Error al escanear la fila:", err)
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error durante el recorrido de filas:", err)
		return nil, err
	}

	return notes, nil
}

func (r *noteRepository) GetByID(id int) (*entities.Note, error) {
	row := r.db.QueryRow("SELECT id, title, description, user_id FROM notes WHERE id = ?", id)
	var note entities.Note

	// Escanear correctamente el resto de los campos
	if err := row.Scan(&note.ID, &note.Title, &note.Description, &note.UserID); err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *noteRepository) Update(note *entities.Note) error {
	query := "UPDATE notes SET title = ?, description = ? WHERE id = ?"
	_, err := r.db.Exec(query, note.Title, note.Description, note.ID)
	return err
}

func (r *noteRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id = ?", id)
	return err
}
