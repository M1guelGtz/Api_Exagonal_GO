package application

import "demob/src/products/domain"

type UpdateProductUseCase struct {
    repository domain.Interface_Product
}

func NewUpdateProductUseCase(repository domain.Interface_Product) *UpdateProductUseCase {
    return &UpdateProductUseCase{repository: repository}
}

func (uc *UpdateProductUseCase) Run(product *domain.Product) error {
    return uc.repository.Update(product)
}
