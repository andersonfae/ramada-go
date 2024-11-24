package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "user:password@tcp(127.0.0.1:3306)/produtos_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}
}
