package application

import (
    "demob/src/products/domain"
)
type DeleteProductUseCase struct {
    repository domain.Interface_Product
}

func NewDeleteProductUseCase(repository domain.Interface_Product) *DeleteProductUseCase {
    return &DeleteProductUseCase{repository: repository}
}

func (uc *DeleteProductUseCase) Run(productId int32) error{
    return uc.repository.Delete(productId)
}