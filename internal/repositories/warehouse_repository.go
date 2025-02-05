package repositories

import (
	"context"
	"gorm.io/gorm"
	"wh-gin/internal/entities"
)

type WarehouseRepository interface {
	Create(ctx context.Context, warehouse entities.WarehouseLocation) error
	GetAll(ctx context.Context, limit, offset int) ([]entities.WarehouseLocation, int, error)
	GetByID(ctx context.Context, id int) (entities.WarehouseLocation, error)
}

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepository{
		db: db,
	}
}

func (w *warehouseRepository) Create(ctx context.Context, warehouse entities.WarehouseLocation) error {
	return w.db.WithContext(ctx).Create(&warehouse).Error
}

func (w *warehouseRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.WarehouseLocation, int, error) {
	var result []entities.WarehouseLocation
	if err := w.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, 0, err
	}

	count := w.db.WithContext(ctx).Find(&[]entities.WarehouseLocation{}).RowsAffected

	return result, int(count), nil
}

func (w *warehouseRepository) GetByID(ctx context.Context, id int) (entities.WarehouseLocation, error) {
	result := entities.WarehouseLocation{}
	if err := w.db.WithContext(ctx).First(&result, id).Error; err != nil {
		return entities.WarehouseLocation{}, err
	}

	return result, nil
}
