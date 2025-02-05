package usecases

import (
	"context"
	"wh-gin/internal/entities"
	"wh-gin/internal/repositories"
)

type ProductUsecase interface {
	InsertProduct(ctx context.Context, product entities.Product) error
	GetAllProducts(ctx context.Context, limit, offset int) ([]entities.Product, int, error)
	GetDetail(ctx context.Context, id int) (entities.Product, error)
	Update(ctx context.Context, product entities.Product) error
	Delete(ctx context.Context, id int) error
}

type productUsecase struct {
	repo repositories.ProductRepository
}

func NewProductUsecase(repository repositories.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repository,
	}
}

func (p *productUsecase) InsertProduct(ctx context.Context, product entities.Product) error {
	return p.repo.Create(ctx, product)
}

func (p *productUsecase) GetAllProducts(ctx context.Context, limit, offset int) ([]entities.Product, int, error) {
	return p.repo.GetAll(ctx, limit, offset)
}

func (p *productUsecase) GetDetail(ctx context.Context, id int) (entities.Product, error) {
	return p.repo.GetByID(ctx, id)
}

func (p *productUsecase) Update(ctx context.Context, product entities.Product) error {
	return p.repo.UpdateByID(ctx, product)
}

func (p *productUsecase) Delete(ctx context.Context, id int) error {
	return p.repo.DeleteByID(ctx, id)
}
