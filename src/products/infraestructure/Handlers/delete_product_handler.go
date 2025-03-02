package handlers

import (
	"demob/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductHandler struct {
	UseCase *application.DeleteProductUseCase
}

func NewDeleteProductHandler(useCase *application.DeleteProductUseCase) *DeleteProductHandler {
	return &DeleteProductHandler{UseCase: useCase}
}

func (h *DeleteProductHandler) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.UseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
