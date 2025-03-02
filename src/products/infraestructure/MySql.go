package infraestructure

import (
	"demob/src/core"
	"demob/src/products/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// 🔹 Guardar un producto
func (mysql *MySQL) Save (product *domain.Product) error {
	query := `INSERT INTO products (nombre, precio, cantidad) VALUES (?, ?, ?)`
	result, err := mysql.conn.DB.Exec(query, product.Nombre, product.Precio, product.Cantidad)
	if err != nil {
		log.Println("Error insertando producto:", err)
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error obteniendo el ID del nuevo producto:", err)
		return err
	}
	fmt.Println("✅ Nuevo Producto creado con ID:", lastID)
	return nil
}

// 🔹 Obtener producto por ID
func (mysql *MySQL) GetById(id int32) (*domain.Product, error) {
	query := `SELECT id, nombre, precio, cantidad FROM products WHERE id = ?`
	row := mysql.conn.DB.QueryRow(query, id)

	var product domain.Product
	err := row.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad)
	if err != nil {
		log.Println("Error leyendo producto:", err)
		return nil, err
	}
	return &product, nil
}

// 🔹 Obtener todos los productos
func (mysql *MySQL) GetAll() ([]*domain.Product, error) {
	query := `SELECT id, nombre, precio, cantidad FROM products`
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Println("Error consultando productos:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad)
		if err != nil {
			log.Println("Error leyendo fila:", err)
			continue
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error durante la iteración de las filas:", err)
		return nil, err
	}

	// Retornar slice vacío si no hay productos en lugar de nil
	if len(products) == 0 {
		return []*domain.Product{}, nil
	}

	return products, nil
}

// 🔹 Actualizar un producto
func (mysql *MySQL) Update(product *domain.Product) error {
	query := `UPDATE products SET nombre = ?, precio = ?, cantidad = ? WHERE id = ?`
	result, err := mysql.conn.DB.Exec(query, product.Nombre, product.Precio, product.Cantidad, product.Id)
	if err != nil {
		log.Println("Error actualizando producto:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error obteniendo el número de filas afectadas:", err)
		return err
	}
	log.Printf("Número de productos actualizados: %d\n", rowsAffected)
	return nil
}

// 🔹 Eliminar un producto
func (mysql *MySQL) Delete(productID int32) error {
	query := `DELETE FROM products WHERE id = ?`
	result, err := mysql.conn.DB.Exec(query, productID)
	if err != nil {
		log.Println("Error eliminando producto:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error obteniendo el número de filas afectadas:", err)
		return err
	}
	log.Printf("Número de productos eliminados: %d\n", rowsAffected)
	return nil
}
