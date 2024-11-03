package product

import (
	"context"
	"fmt"
	"tech-testing/src/models"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductRepository) FindAll(ctx context.Context, limit, skip int) ([]*models.Product, error) {
	var products []*models.Product
	if err := r.db.WithContext(ctx).Limit(limit).Offset(skip).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Update(ctx context.Context, id uint, product *models.Product) error {
	result := r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(product)
	return result.Error
}

func (r *ProductRepository) Remove(ctx context.Context, id uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Update("deleted_at", now).Error
}

func (r *ProductRepository) FindById(ctx context.Context, id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}
