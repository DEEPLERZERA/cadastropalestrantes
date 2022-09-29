package controllers

import (
	"CadastroPalestrantes/database"
	"CadastroPalestrantes/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"strings"

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

	havestringCPF := strings.ContainsAny(palestrante.CPF, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_+*/!@#$%¨&*()_+}{^~´`][}{><,.;:?/|")
	havestringRG := strings.ContainsAny(palestrante.RG, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_+*/!@#$%¨&*()_+}{^~´`][}{><,.;:?/|")
	havenumber := strings.ContainsAny(palestrante.Nome, "0123456789")
	cpfexists := database.DB.Where(&models.Palestrante{CPF: palestrante.CPF}).First(&palestrante)  //Pode usar pra verificar cpf
	rgexists := database.DB.Where(&models.Palestrante{RG: palestrante.RG}).First(&palestrante)  //Pode usar pra verificar rg


	if cpfexists.RowsAffected == 1 {
		c.JSON(404, gin.H{"error": "CPF já cadastrado"})
		return
	}

	if rgexists.RowsAffected == 1 {
		c.JSON(404, gin.H{"error": "RG já cadastrado"})
		return
	}

	if len(palestrante.CPF) != 11 {
		c.JSON(404, gin.H{"error": "CPF deve ter 11 dígitos"})
		return
	}

	if len(palestrante.RG) != 9 {
		c.JSON(404, gin.H{"error": "RG deve ter 9 dígitos"})
		return
	}
	

	if palestrante.CPF == "" || palestrante.Nome == "" || palestrante.RG == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados incompletos"})
		return
	}
	if havestringCPF {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro CPF deve ser número obrigatoriamente!"})
		return
	}else if havestringRG {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro RG deve ser número obrigatoriamente!"})
		return
	}else if havenumber {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro nome só pode possuir caracteres, Números não são aceitos!"})
		return
	}else {
		database.DB.Create(&palestrante)
		c.JSON(200, palestrante)
	}
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

	havestring := strings.ContainsAny(id, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_+*/!@#$%¨&*()_+}{^~´`][}{><,.;:?/|")
	iduint64, err:= strconv.ParseUint(id, 10, 32) 
	iduint := uint(iduint64)
	if err != nil {
		c.JSON(404, gin.H{"error": "Erro ao deletar"})
		return
	}
	idexists  := database.DB.Where(gorm.Model{ID: iduint}).First(&palestrante)  //Pode usar pra verificar id

	

	if idexists.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "ID não encontrado"})
		return
	}

	if havestring {
		c.JSON(404, gin.H{"error": "ID deve ser número obrigatoriamente!"})
		return
	} else {
		c.JSON(200, gin.H{"message": "Palestrante deletado com sucesso"})
		database.DB.Delete(&palestrante, id)
	}
	
}


func BuscaPalestrantePorCPF(c *gin.Context) {
	var palestrante models.Palestrante

	cpf := c.Param("cpf")
	havestring := strings.ContainsAny(cpf, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_+*/!@#$%¨&*()_+}{^~´`][}{><,.;:?/|")
	database.DB.Where(&models.Palestrante{CPF: cpf}).First(&palestrante)  //Pode usar pra verificar cpf 

	if palestrante.ID == 0 {
		c.JSON(404, gin.H{"error": "Palestrante não encontrado"})
		return
	} 
	
	if havestring {
		c.JSON(404, gin.H{"error": "CPF deve ser número obrigatoriamente!"})
		return
	} else {
		c.JSON(200, palestrante)
	}
	
}