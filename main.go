package main

import (
	"demob/src/application"
	"demob/src/infraestructure"
	"fmt"
	"log"
	//"os"
	"github.com/gin-gonic/gin"
)

func main() {
	/*
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	*/
	dsn := fmt.Sprintf("root:passw0rd@tcp(127.0.0.1:8080)/Demo?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := infraestructure.NewMySql(dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	createProductUseCase := application.NewCreateProductUseCase(db)
	viewAllProductsUseCase := application.NewViewAllProductsUseCase(db)
	updateProductuseCase := application.NewUpdateProductUseCase(db)
	deleteProductUseCase := application.NewDeleteProductUseCase(db)
	productController := infraestructure.NewProductController(createProductUseCase, viewAllProductsUseCase, updateProductuseCase, deleteProductUseCase)

	r := gin.Default()
	infraestructure.RegisterProductRoutes(r, productController)

	if err := r.Run(":" + "3000"); err != nil {
		panic(err)
	}
}
