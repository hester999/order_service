package repo

import (
	"context"
	"order/internal/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderByOrderUID(ctx context.Context, id string) (model.Order, error)
}

type OrderPreload interface {
	GetAllOrders(ctx context.Context) ([]model.Order, error)
}

type ItemRepository interface {
	CreateItem(ctx context.Context, item model.Item) (model.Item, error)
	GetItemsByOrderUID(ctx context.Context, orderUID string) ([]model.Item, error)
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment model.Payment) (model.Payment, error)
	GetPaymentsByOrderUID(ctx context.Context, orderUID string) (model.Payment, error)
}

type DeliveryRepository interface {
	CreateDelivery(ctx context.Context, delivery model.Delivery) (model.Delivery, error)
	GetDeliveryByOrderUID(ctx context.Context, deliveryID string) (model.Delivery, error)
}
