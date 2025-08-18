package postgress

import (
	"order/internal/repo"
	"order/internal/repo/postgress/delivery"
	"order/internal/repo/postgress/item"
	"order/internal/repo/postgress/order"
	"order/internal/repo/postgress/payment"

	"github.com/jmoiron/sqlx"
)

type PostgresStorage struct {
	repo.OrderRepository
	repo.ItemRepository
	repo.DeliveryRepository
	repo.PaymentRepository
	repo.OrderPreload
}

func NewPostgresStorage(db *sqlx.DB) *PostgresStorage {
	return &PostgresStorage{
		OrderRepository:    order.NewOrderRepo(db),
		ItemRepository:     item.NewItemRepo(db),
		DeliveryRepository: delivery.NewDeliveryRepo(db),
		PaymentRepository:  payment.NewPaymentRepo(db),
		OrderPreload:       order.NewOrderRepo(db),
	}
}
