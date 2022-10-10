package routes

import (
	"tfalc/Api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/", controllers.ExiberPaginaInicial)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Greetings)
	r.POST("/alunos", controllers.CriaNovoAluno)

	r.Run()
}
