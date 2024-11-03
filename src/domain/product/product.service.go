package product

import (
	"context"
	"tech-testing/src/models"
)

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) GetAllProducts(ctx context.Context, limit, skip int) ([]*models.Product, error) {
	return s.repo.FindAll(ctx, limit, skip)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uint, product *models.Product) error {
	return s.repo.Update(ctx, id, product)
}

func (s *ProductService) RemoveProduct(ctx context.Context, id uint) error {
	return s.repo.Remove(ctx, id)
}

func (s *ProductService) FindProductById(ctx context.Context, id uint) (*models.Product, error) {
	return s.repo.FindById(ctx, id)
}
