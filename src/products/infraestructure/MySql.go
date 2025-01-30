package infraestructure

import (
	"demob/src/products/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	db *gorm.DB
}

func NewMySql(dsn string) (*MySql, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// No realizar auto-migración automática, ya que la tabla puede estar correctamente configurada
	// if err := db.AutoMigrate(&domain.Product{}); err != nil {
	//     return nil, err
	// }

	return &MySql{db: db}, nil
}

func (mysql *MySql) Save(product *domain.Product) error {
	result := mysql.db.Create(product)
	return result.Error
}

func (mysql *MySql) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	result := mysql.db.Find(&products)
	return products, result.Error
}

func (mysql *MySql) Update(product *domain.Product) error {
	result := mysql.db.Save(product)
	return result.Error
}

func (mysql *MySql) Delete(productID int32) error {
	result := mysql.db.Delete(&domain.Product{}, productID)
	return result.Error
}
