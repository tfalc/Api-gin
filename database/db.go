package database

import (
	"log"
	"tfalc/Api-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBanco() {
	conexao := "user=postgres dbname=tfalc_alunos password=123 host=localhost sslmode=disable"
	DB, err := gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	}

	DB.AutoMigrate(&models.Aluno{})
}
