package products

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
func (p *PostgresRepository) GetAll(ctx context.Context) ([]Products, error) {
	var products []Products
	rows, err := p.db.QueryContext(ctx, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Products
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
func (p *PostgresRepository) GetProduct(ctx context.Context, id int) (*Products, error) {
	var product Products
	err := p.db.QueryRowContext(ctx, "SELECT * FROM products WHERE id = $1", id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}
func (p *PostgresRepository) Search(ctx context.Context, query string) ([]Products, error) {
	var products []Products
	rows, err := p.db.Query(`
		SELECT *
		FROM products
		WHERE name ILIKE '%' || $1 || '%'
		   OR description ILIKE '%' || $1 || '%'
		ORDER BY created_at DESC
	`, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Products
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("Отдаем репо")
	return products, nil
}
