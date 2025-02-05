package repositories

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"wh-gin/internal/entities"
)

type ProductRepository interface {
	Create(ctx context.Context, product entities.Product) error
	GetAll(ctx context.Context, limit, offset int) ([]entities.Product, int, error)
	GetByID(ctx context.Context, id int) (entities.Product, error)
	UpdateByID(ctx context.Context, product entities.Product) error
	DeleteByID(ctx context.Context, id int) error
	UpdateQuantity(ctx context.Context, id, quantity int) error
	GetCurrentStock(ctx context.Context, warehouseID int) (int, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) Create(ctx context.Context, product entities.Product) error {
	if err := p.db.WithContext(ctx).Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p *productRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Product, int, error) {
	var result []entities.Product
	if err := p.db.WithContext(ctx).Preload("WarehouseLocation").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, 0, err
	}

	count := p.db.WithContext(ctx).Find(&[]entities.Product{}).RowsAffected

	return result, int(count), nil
}

func (p *productRepository) GetByID(ctx context.Context, id int) (entities.Product, error) {
	var result entities.Product
	if err := p.db.WithContext(ctx).Preload("WarehouseLocation").First(&result, "id = ?", id).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (p *productRepository) UpdateByID(ctx context.Context, product entities.Product) error {
	if err := p.db.WithContext(ctx).Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p *productRepository) DeleteByID(ctx context.Context, id int) error {
	if err := p.db.WithContext(ctx).Delete(&entities.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (p *productRepository) UpdateQuantity(ctx context.Context, id, quantity int) error {
	if err := p.db.WithContext(ctx).Model(&entities.Product{}).Where("id = ?", id).Update("quantity", quantity).Error; err != nil {
		return err
	}
	return nil
}

func (p *productRepository) GetCurrentStock(ctx context.Context, warehouseID int) (int, error) {
	var currentStock int

	err := p.db.WithContext(ctx).Model(&entities.Product{}).
		Select("COALESCE(SUM(quantity), 0)").
		Where("location_id = ?", warehouseID).
		Scan(&currentStock).Error

	if err != nil {
		return 0, fmt.Errorf("failed to get current stock: %w", err)
	}

	return currentStock, nil
}
