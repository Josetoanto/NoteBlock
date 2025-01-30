package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // Importa el dialecto de MySQL
    "log"
)

var DB *gorm.DB


func Connect() (*gorm.DB, error) {
    dsn := "root:tono1234@/apirestgo?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
        return nil, err
    }
    DB = db // Asignar la conexión a la variable exportada DB
    return db, nil
}
