package application

import "demob/src/domain"

type CreateProductUseCase struct {
	repository domain.IProduct
}

func NewCreateProductUseCase(repository domain.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{repository: repository}
}

func (uc *CreateProductUseCase) Run(name string, price float32, cantidad float32) error {
	product := domain.NewProduct(name, price, cantidad)
	return uc.repository.Save(product)
}
