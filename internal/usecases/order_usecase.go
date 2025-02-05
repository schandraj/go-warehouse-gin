package usecases

import (
	"context"
	"fmt"
	"time"
	"wh-gin/internal/entities"
	"wh-gin/internal/repositories"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	GetAllOrders(ctx context.Context, limit, offset int) ([]entities.Order, int, error)
	GetByID(ctx context.Context, id int) (entities.Order, error)
	ExecuteOrder(ctx context.Context)
}

type orderUsecase struct {
	repo     repositories.OrderRepository
	prodRepo repositories.ProductRepository
	whRepo   repositories.WarehouseRepository
}

func NewOrderUsecase(repo repositories.OrderRepository, repository repositories.ProductRepository, warehouseRepository repositories.WarehouseRepository) OrderUsecase {
	return &orderUsecase{
		repo:     repo,
		prodRepo: repository,
		whRepo:   warehouseRepository,
	}
}

func (o *orderUsecase) CreateOrder(ctx context.Context, order entities.Order) error {
	return o.repo.CreateOrder(ctx, order)
}

func (o *orderUsecase) GetAllOrders(ctx context.Context, limit, offset int) ([]entities.Order, int, error) {
	return o.repo.GetAll(ctx, limit, offset)
}

func (o *orderUsecase) GetByID(ctx context.Context, id int) (entities.Order, error) {
	return o.repo.GetByID(ctx, id)
}

func (o *orderUsecase) ExecuteOrder(ctx context.Context) {
	for {
		orders, err := o.repo.GetPendingOrders(ctx)
		if err != nil {
			fmt.Println("Error fetching orders:", err) // Log the error instead of returning
			time.Sleep(2 * time.Second)                // Prevent tight looping on failure
			continue
		}

		for _, order := range orders {
			switch order.Type {
			case "receive":
				fmt.Println("Executing order with order id : ", order.ID)
				prod, errProd := o.prodRepo.GetByID(ctx, int(order.ProductID))
				if errProd != nil {
					fmt.Println("Error fetching product:", errProd)
					continue
				}

				stock, errStock := o.prodRepo.GetCurrentStock(ctx, int(prod.LocationID))
				if errStock != nil {
					fmt.Println("Error fetching stock:", errStock)
					continue
				}

				wh, errWh := o.whRepo.GetByID(ctx, int(prod.LocationID))
				if errWh != nil {
					fmt.Println("Error fetching wh:", errWh)
					continue
				}

				quantity := prod.Quantity + order.Quantity

				if wh.Capacity < stock+order.Quantity {
					fmt.Printf("Warehouse Capacity is not enough")
					continue
				}

				if err = o.prodRepo.UpdateQuantity(ctx, int(prod.ID), quantity); err != nil {
					fmt.Println("Error updating quantity:", err)
					continue
				}

				if err = o.repo.UpdateStatus(ctx, int(order.ID), "completed"); err != nil {
					fmt.Println("Error updating order status:", err)
					continue
				}

			case "ship":
				fmt.Println("Executing order with order id : ", order.ID)
				prod, errProd := o.prodRepo.GetByID(ctx, int(order.ProductID))
				if errProd != nil {
					fmt.Println("Error fetching product:", errProd)
					continue
				}

				if prod.Quantity < order.Quantity {
					fmt.Println("Insufficient stock for order:", order.ID)
					continue
				}

				quantity := prod.Quantity - order.Quantity
				if err = o.prodRepo.UpdateQuantity(ctx, int(prod.ID), quantity); err != nil {
					fmt.Println("Error updating quantity:", err)
					continue
				}

				if err = o.repo.UpdateStatus(ctx, int(order.ID), "completed"); err != nil {
					fmt.Println("Error updating order status:", err)
					continue
				}
			}
		}

		time.Sleep(5 * time.Second)
	}
}
