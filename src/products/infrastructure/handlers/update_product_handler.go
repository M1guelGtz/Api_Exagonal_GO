package handlers

import (
	"demob/src/products/application"
	"demob/src/products/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductHandler struct {
	UseCase *application.UpdateProductUseCase
}

func NewUpdateProductHandler(useCase *application.UpdateProductUseCase) *UpdateProductHandler {
	return &UpdateProductHandler{UseCase: useCase}
}

func (pc *UpdateProductHandler) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Id = int32(id)

	if err := pc.UseCase.Run(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
