package handlers

import (
	"demob/src/products/application"
	"demob/src/products/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductHandler struct {
	UseCase *application.CreateProductUseCase
}

func NewCreateProductHandler(useCase *application.CreateProductUseCase) *CreateProductHandler {
	return &CreateProductHandler{UseCase: useCase}
}

func (h *CreateProductHandler) Handle(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.Run(product.Nombre, product.Precio, product.Cantidad); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Producto creado correctamente",
		"product": gin.H{
			"nombre":   &product.Nombre,
			"cantidad": &product.Cantidad,
			"precio":   &product.Precio,
		},
	})
}
