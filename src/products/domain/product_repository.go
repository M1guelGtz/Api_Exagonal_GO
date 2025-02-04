package domain

type Interface_Product interface {
	Save(product *Product) error
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(productId int32) error
	GetById(id int32) (*Product, error)
}