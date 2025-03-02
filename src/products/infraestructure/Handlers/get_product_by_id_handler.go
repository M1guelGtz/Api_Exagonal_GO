package handlers

import (
	"demob/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIdHandler struct {
	UseCase *application.ViewProdByIdUseCase
}

func NewGetProductByIdHandler(useCase *application.ViewProdByIdUseCase) *GetProductByIdHandler {
	return &GetProductByIdHandler{UseCase: useCase}
}

func (h *GetProductByIdHandler) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.UseCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
