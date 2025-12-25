package products

import "context"

type Repository interface {
	GetProduct(ctx context.Context, id int) (*Products, error)
	GetAll(ctx context.Context) ([]Products, error)
	Search(ctx context.Context, query string) ([]Products, error)
}
