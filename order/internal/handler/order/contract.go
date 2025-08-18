package order

import (
	"context"
	"order/internal/model"
)

type Order interface {
	GetOrder(ctx context.Context, id string) (model.Order, error)
}
