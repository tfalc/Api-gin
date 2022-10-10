package main

import (
	"tfalc/Api-gin/database"
	"tfalc/Api-gin/models"
	"tfalc/Api-gin/routes"
)

func main() {
	database.ConectaComBanco()
	models.Alunos = []models.Aluno{
		{Nome: "Thiago", CPF: "00000000000", RG: "1234567891"},
		{Nome: "Janaina", CPF: "00000001234", RG: "4321567891"},
	}
	routes.HandleRequests()
}
