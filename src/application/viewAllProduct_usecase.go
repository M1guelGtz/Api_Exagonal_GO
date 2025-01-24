package application

import "demob/src/domain"

type ViewAllProductsUseCase struct {
    repository domain.IProduct
}

func NewViewAllProductsUseCase(repository domain.IProduct) *ViewAllProductsUseCase {
    return &ViewAllProductsUseCase{repository: repository}
}

func (uc *ViewAllProductsUseCase) Run() ([]*domain.Product, error) {
    return uc.repository.GetAll()
}
