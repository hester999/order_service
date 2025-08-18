package order

import "order/internal/repo"

type Storage interface {
	repo.OrderRepository
	repo.ItemRepository
	repo.DeliveryRepository
	repo.PaymentRepository
	repo.OrderPreload
}

type CacheStorage interface {
	repo.OrderRepository
}
