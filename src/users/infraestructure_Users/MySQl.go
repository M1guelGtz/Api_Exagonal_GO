package infraestructureusers

import (
	domainusers "demob/src/users/domain_users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL implementa la interfaz UserInterface
type MySQL struct {
	db *gorm.DB
}

// NewMySQL inicializa la conexión a la base de datos
func NewMySQL(dsn string) (*MySQL, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySQL{db: db}, nil
}

// Implementación de la interfaz UserInterface
func (m *MySQL) Create(user *domainusers.User) error {
	return m.db.Create(user).Error
}

func (m *MySQL) GetUsers() ([]*domainusers.User, error) {
	var users []*domainusers.User
	err := m.db.Find(&users).Error
	return users, err
}

func (m *MySQL) GetUserById(id int32) (*domainusers.User, error) {
	var user domainusers.User
	err := m.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (m *MySQL) UpdateUser(user *domainusers.User) error {
	return m.db.Save(user).Error
}

func (m *MySQL) DeleteUser(id int32) error {
	return m.db.Delete(&domainusers.User{}, id).Error
}
