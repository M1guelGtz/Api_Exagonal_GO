package domain

type Product struct {
	Id       int32   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre   string  `gorm:"size:255" json:"nombre"`
	Precio   float32 `gorm:"type:decimal(10,2)" json:"precio"`
	Cantidad float32 `gorm:"type:decimal(10,2)" json:"cantidad"`
}
func NewProduct(name string, price float32, cantidad float32) *Product {
	return &Product{Nombre: name, Precio: price, Cantidad: cantidad}
}

func (p *Product) GetName() string {
	return p.Nombre
}

func (p *Product) SetName(name string) {
	p.Nombre = name
}
