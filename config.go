package main

import (
    "demob/src/application"
    "demob/src/infraestructure"
    "github.com/gin-gonic/gin"
)

func Setup() (*gin.Engine, error) {
    dsn := "root:root@tcp(127.0.0.1:3306)/databasegofirst?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := infraestructure.NewMySql(dsn)
    if err != nil {
        return nil, err
    }

    createProductUseCase := application.NewCreateProductUseCase(db)
    viewAllProductsUseCase := application.NewViewAllProductsUseCase(db)
    updateProductUseCase := application.NewUpdateProductUseCase(db)
    deleteProductUseCase := application.NewDeleteProductUseCase(db)
    productController := infraestructure.NewProductController(createProductUseCase, viewAllProductsUseCase, updateProductUseCase, deleteProductUseCase)

    r := gin.Default()
    infraestructure.RegisterProductRoutes(r, productController)

    return r, nil
}
