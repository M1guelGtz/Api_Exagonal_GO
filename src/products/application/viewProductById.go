package application
import "demob/src/products/domain"

type ViewProdByIdUseCase struct {
	repository domain.Interface_Product
}

func NewViewPrByIdUseCase(repository domain.Interface_Product) *ViewProdByIdUseCase {
	return &ViewProdByIdUseCase{repository: repository}
}

func (uc *ViewProdByIdUseCase) Execute(id int32) (*domain.Product, error) {
	return uc.repository.GetById(id)
}
