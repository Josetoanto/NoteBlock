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
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   *time.Time `json:"created_at"`
}

func main() {
	db, err := sql.Open("mysql", "root:tono1234@tcp(localhost:3306)/notes_db")
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)
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
	rows, err := db.Query("SELECT id, username, password, created_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("Error en la consulta: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var createdAtRaw interface{}

		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &createdAtRaw); err != nil {
			return nil, fmt.Errorf("Error al leer los datos del usuario: %v", err)
		}

		if createdAtRaw != nil {
			createdAt, ok := createdAtRaw.(time.Time)
			if ok {
				user.CreatedAt = &createdAt
			}
		}

		users = append(users, user)
	}
	return users, nil
}
