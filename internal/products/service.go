package products

import (
	"context"
	"errors"
)

type Service interface {
	Product(ctx context.Context, id int) (*Products, error)
	Products(ctx context.Context) ([]Products, error)
	Search(ctx context.Context, query string) ([]Products, error)
}

var ErrId = errors.New("ID 0")
var ErrInputNull = errors.New("Input is null")

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Product(ctx context.Context, id int) (*Products, error) {
	if id < 1 {
		return nil, ErrId
	}
	return s.repo.GetProduct(ctx, id)
}
func (s *service) Products(ctx context.Context) ([]Products, error) {
	return s.repo.GetAll(ctx)
}
func (s *service) Search(ctx context.Context, query string) ([]Products, error) {
	if query == "" {
		return nil, ErrInputNull
	}
	return s.repo.Search(ctx, query)
}
