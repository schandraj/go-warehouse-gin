package repositories

import (
	"context"
	"gorm.io/gorm"
	"wh-gin/internal/entities"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	GetAll(ctx context.Context, limit, offset int) ([]entities.Order, int, error)
	GetByID(ctx context.Context, id int) (entities.Order, error)
	GetPendingOrders(ctx context.Context) ([]entities.Order, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) CreateOrder(ctx context.Context, order entities.Order) error {
	return o.db.WithContext(ctx).Create(&order).Error
}

func (o *orderRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Order, int, error) {
	var result []entities.Order
	if err := o.db.WithContext(ctx).Preload("Product.WarehouseLocation").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, 0, err
	}

	count := o.db.WithContext(ctx).Find(&[]entities.Order{}).RowsAffected

	return result, int(count), nil
}

func (o *orderRepository) GetByID(ctx context.Context, id int) (entities.Order, error) {
	var result entities.Order
	if err := o.db.WithContext(ctx).Preload("Product.WarehouseLocation").First(&result, id).Error; err != nil {
		return entities.Order{}, err
	}
	return result, nil
}

func (o *orderRepository) GetPendingOrders(ctx context.Context) ([]entities.Order, error) {
	var result []entities.Order
	if err := o.db.WithContext(ctx).Where("status = ?", "pending").Limit(10).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (o *orderRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	if err := o.db.WithContext(ctx).Model(&entities.Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
