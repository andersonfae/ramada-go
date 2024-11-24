package routes

import (
	"api-produtos/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/produtos", controllers.ListarProdutos)
	r.GET("/produtos/:id", controllers.BuscarProdutoPorID)
	r.POST("/produtos", controllers.AdicionarProduto)
	// Demais rotas

	return r
}
