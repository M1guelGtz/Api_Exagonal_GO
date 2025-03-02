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


func NewMySQL() (*MySQL) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err);
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(product *domain.Product) error {
	query := ``
	result, err := mysql.conn.DB.Exec(query)
	if err != nil {
		log.Println("Error insertando mensaje:", err)
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error obteniendo el ID del nuevo mensaje:", err)
		return err
	}
	fmt.Println("✅ Nuevo Mensaje creado con ID:", lastID)
	return nil
}

func (mysql *MySQL) GetById(id int32) (*domain.Product, error) {
	query := fmt.Sprintf(`SELECT id, nombre, precio, cantidad FROM productos WHERE id = %d`, id)
    row := mysql.conn.DB.QueryRow(query)
    var product domain.Product
    err := row.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad)
    if err != nil {
        log.Println("Error leyendo producto:", err)
        return nil, err
    }
    return &product, nil
}

func (mysql *MySQL) GetAll() ([]*domain.Product, error) {
	query := ``
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
        log.Println("Error consultando productos:", err)
        return nil, err
    }
	defer rows.Close()
	var products []*domain.Product
	count := 0
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Nombre, &product.Precio, &product.Cantidad)
		if err != nil {
            log.Println("Error leyendo fila:", err)
            continue
        }
		products = append(products, &product)
		count++
	}
	if err := rows.Err(); err != nil {
		log.Println("Error durante la iteración de las filas:", err)
		return nil, err
	}
	log.Printf("numero de productos obtenidos: %d\n", count)
	log.Fatalf("productos obtenidos: %+v\n", products)
	return products, nil
}

func (mysql *MySQL) Update(product *domain.Product) error {
	query := "UPDATE"
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

func (mysql *MySQL) Delete(productID int32) error {
	query := "DELETE FROM"
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
