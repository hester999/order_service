package consumer

import (
	"context"
	"order/internal/model"
)

type Order interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
}

type MessageReader interface {
	ReadMessage(ctx context.Context) ([]byte, error)
	Close() error
}
