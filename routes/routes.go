package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tfalc/Api-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/", controllers.ExiberPaginaInicial)
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	r.Run()
}
