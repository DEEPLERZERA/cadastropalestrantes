package controllers 

import (
	"CadastroPalestrantes/database"
	"CadastroPalestrantes/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosPalestrantes(c *gin.Context) {
	var palestrantes []models.Palestrante
	database.DB.Find(&palestrantes)
	c.JSON(200, palestrantes)
}

func CriaNovoPalestrante(c *gin.Context) {
	var palestrante models.Palestrante
	if err := c.ShouldBindJSON(&palestrante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&palestrante)
	c.JSON(200, palestrante)
}
 
func ExibePalestrantPorID(c *gin.Context) {
	var palestrante models.Palestrante
	id := c.Params.ByName("id")
	database.DB.First(&palestrante, id)

	if palestrante.ID == 0 {
		c.JSON(404, gin.H{"error": "Palestrante não encontrado"})
		return
	}

	c.JSON(200, palestrante)

}

func DeletaPalestrante(c *gin.Context) {
	var palestrante models.Palestrante
	id := c.Params.ByName("id")
	database.DB.Delete(&palestrante, id)

	if palestrante.ID == 0 {
		c.JSON(404, gin.H{"error": "Palestrante não encontrado"})
		return
	}

	c.JSON(200, gin.H{"data": "Palestrante deletado com sucesso"})
}


func BuscaPalestrantePorCPF(c *gin.Context) {
	var palestrante models.Palestrante

	cpf := c.Param("cpf")
	database.DB.Where(&models.Palestrante{CPF: cpf}).First(&palestrante)

	if palestrante.ID == 0 {
		c.JSON(404, gin.H{"error": "Palestrante não encontrado"})
		return
	}
	c.JSON(200, palestrante)
}