package controllers

import "net/http"

func ExiberPaginaInicial(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html> <body> <h1>Página Inicial</h1> </body> </html>"))
}

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Thiago",
	})
}
