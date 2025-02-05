package usecases

import (
	"context"
	"wh-gin/internal/entities"
	"wh-gin/internal/repositories"
)

type WarehouseUsecase interface {
	CreateWarehouse(ctx context.Context, warehouse entities.WarehouseLocation) error
	GetAll(ctx context.Context, limit, offset int) ([]entities.WarehouseLocation, int, error)
}

type warehouseUsecase struct {
	repo repositories.WarehouseRepository
}

func NewWarehouseUsecase(repository repositories.WarehouseRepository) WarehouseUsecase {
	return &warehouseUsecase{
		repo: repository,
	}
}

func (w *warehouseUsecase) CreateWarehouse(ctx context.Context, warehouse entities.WarehouseLocation) error {
	return w.repo.Create(ctx, warehouse)
}

func (w *warehouseUsecase) GetAll(ctx context.Context, limit, offset int) ([]entities.WarehouseLocation, int, error) {
	return w.repo.GetAll(ctx, limit, offset)
}
