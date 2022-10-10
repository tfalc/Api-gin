package controllers

import (
	"net/http"

	"tfalc/Api-gin/database"
	"tfalc/Api-gin/models"

	"github.com/gin-gonic/gin"
)

func ExiberPaginaInicial(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html> <body> <h1>Página Inicial</h1> </body> </html>"))
}

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Greetings(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API DIZ: ": "Olá, " + nome + "!",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
