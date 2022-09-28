package routes

import (
	"CadastroPalestrantes/controllers"

	"github.com/gin-gonic/gin"

)

func HandleRequests() {
	r := gin.Default()

	r.GET("/palestrantes", controllers.ExibeTodosPalestrantes)
	r.POST("/palestrantes", controllers.CriaNovoPalestrante)
	r.GET("/palestrantes/:id", controllers.ExibePalestrantPorID)
	r.DELETE("/palestrantes/:id", controllers.DeletaPalestrante)
	r.GET("/palestrantes/cpf/:cpf", controllers.BuscaPalestrantePorCPF)
	r.Run()
	
}