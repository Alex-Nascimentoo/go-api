package repository

import (
	"database/sql"
	"fmt"

	"github.com/alex-nascimentoo/go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var productList []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			fmt.Println(err)
			return nil, err
		}

		productList = append(productList, product)
	}

	return productList, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query := "SELECT * FROM product WHERE id = $1"
	row := pr.connection.QueryRow(query, id)

	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"
	var id int
	err := pr.connection.QueryRow(query, product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
