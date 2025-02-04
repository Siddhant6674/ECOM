package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Siddhant6674/ECOM/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM Products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {

		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p) // Dereference the pointer before appending
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Store) GetProductsByIDs(ProductIDs []int) ([]types.Product, error) {
	placeholder := strings.Repeat("?,", len(ProductIDs))
	placeholder = placeholder[:len(placeholder)-1] // Remove the trailing comma

	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", placeholder)

	//convert ProductIDs into []interface{}
	args := make([]interface{}, len(ProductIDs))
	for i, v := range ProductIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	Products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		Products = append(Products, *p)
	}
	return Products, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	_, err :=
		s.db.Exec("INSERT INTO products(name,description,image,price,quantity)VALUES(?,?,?,?,?)",
			product.Name,
			product.Description,
			product.Image,
			product.Price,
			product.Quantity,
		)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?,price=?,image=?,description=?,quantity=? WHERE id=?",
		product.Name,
		product.Price,
		product.Image,
		product.Description,
		product.Quantity,
		product.ID)
	if err != nil {
		return err
	}
	return nil
}
