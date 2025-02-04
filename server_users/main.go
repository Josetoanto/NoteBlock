package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
}

func main() {
	db, err := sql.Open("mysql", "root:tono1234@tcp(localhost:3306)/notes_db")
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)
	} else {
		log.Println("Conexión exitosa a la base de datos")
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/users/short", func(c *gin.Context) {
		users, err := getAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.GET("/users/long", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(c.Writer)

		users, err := getAllUsers(db)
		if err != nil {
			log.Println("Error obteniendo usuarios:", err)
			return
		}
		for _, user := range users {
			time.Sleep(2 * time.Second)
			if err := encoder.Encode(user); err != nil {
				log.Println("Error enviando usuario:", err)
				return
			}
			c.Writer.Flush()
		}
	})

	fmt.Println("Servidor de usuarios corriendo en :8081")
	r.Run(":8081")
}

func getAllUsers(db *sql.DB) ([]User, error) {
	// Modificamos la consulta para no incluir 'created_at'
	rows, err := db.Query("SELECT id, title, description, user_id FROM users")
	if err != nil {
		return nil, fmt.Errorf("Error en la consulta: %v", err)
	}
	defer rows.Close()

	var users []User
	count := 0
	for rows.Next() {
		var user User
		// No tomamos en cuenta 'created_at', y solo leemos los demás campos
		if err := rows.Scan(&user.ID, &user.Title, &user.Description, &user.UserID); err != nil {
			return nil, fmt.Errorf("Error al leer los datos del usuario: %v", err)
		}
		// Si tienes datos para 'created_at', puedes asignarlo aquí, si no lo deseas
		// user.CreatedAt = time.Time{} // Opcional si quieres dejar 'CreatedAt' como cero
		users = append(users, user)
		count++
	}

	if count == 0 {
		log.Println("No se encontraron usuarios en la base de datos.")
	}

	return users, nil
}
