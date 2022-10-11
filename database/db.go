package database

import (
	"log"
	"tfalc/Api-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func ConectaComBanco() {
	conexao := "host=localhost user=postgres password=123 dbname=tfalc_alunos port=5432 sslmode=disable"
	DB, Err := gorm.Open(postgres.Open(conexao), &gorm.Config{})
	if Err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	}

	DB.AutoMigrate(&models.Aluno{})
}
