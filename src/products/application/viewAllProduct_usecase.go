package application

import "demob/src/products/domain"

type ViewAllProductsUseCase struct {
    repository domain.Interface_Product
}

func NewViewAllProductsUseCase(repository domain.Interface_Product) *ViewAllProductsUseCase {
    return &ViewAllProductsUseCase{repository: repository}
}

func (uc *ViewAllProductsUseCase) Run() ([]*domain.Product, error) {
    return uc.repository.GetAll()
}
