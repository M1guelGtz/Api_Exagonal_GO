package main

import (
	//"os"
	"demob/src/products/infraestructure"

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
	//dsn := ("root:passw0rd@tcp(127.0.0.1:8080)/Demo?charset=utf8mb4&parseTime=True&loc=Local")

	//productController := infraestructure.NewProductController(createProductUseCase, viewAllProductsUseCase, updateProductuseCase, deleteProductUseCase, viewProductById)
	r := gin.Default()
	infraestructure.Init(r)
	if err := r.Run(":"+"3000"); err != nil {
		panic(err)
	}

	// Usar el middleware de CORS
}
