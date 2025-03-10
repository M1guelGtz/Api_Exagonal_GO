package handlers

import (
	"demob/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductsHandler struct {
	UseCase *application.ViewAllProductsUseCase
}

func NewGetAllProductsHandler(useCase *application.ViewAllProductsUseCase) *GetAllProductsHandler {
	return &GetAllProductsHandler{UseCase: useCase}
}

func (h *GetAllProductsHandler) Handle(c *gin.Context) {
	products, err := h.UseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
