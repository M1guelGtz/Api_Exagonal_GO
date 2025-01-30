package application

import "demob/src/products/domain"

type CreateProductUseCase struct {
	repository domain.Interface_Product
}

func NewCreateProductUseCase(repository domain.Interface_Product) *CreateProductUseCase {
	return &CreateProductUseCase{repository: repository}
}

func (uc *CreateProductUseCase) Run(name string, price float32, cantidad float32) error {
	product := domain.NewProduct(name, price, cantidad)
	return uc.repository.Save(product)
}
