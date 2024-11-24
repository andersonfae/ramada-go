package controllers

import (
	"api-produtos/database"
	"api-produtos/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListarProdutos(c *gin.Context) {
	var produtos []models.Produto
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10
	offset := (page - 1) * limit

	database.DB.Limit(limit).Offset(offset).Find(&produtos)
	c.JSON(http.StatusOK, produtos)
}

func BuscarProdutoPorID(c *gin.Context) {
	id := c.Param("id")
	var produto models.Produto

	if err := database.DB.First(&produto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	c.JSON(http.StatusOK, produto)
}

func AdicionarProduto(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&produto)
	c.JSON(http.StatusCreated, produto)
}
