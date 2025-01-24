package domain

type Usuarios struct {
	Id       int64  `gorm: "primaryKey;autoIncrement" json:"id"`
	Nombre   string `gorm: "size:255" json: "nombre"`
	Apellido string `gorm: "size: 255" json: "apellido"`
	Email    string `gorm: "uniqueIndex; size:255" json: "email"`
	Password string `gorm: "size: 255" json: "password"`
}
