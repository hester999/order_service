package order

import (
	"context"
	"log"
	"order/internal/model"

	"github.com/google/uuid"
)

type FullOrder struct {
	db    Storage
	cache CacheStorage
}

func NewFullOrder(db Storage, cache CacheStorage) *FullOrder {
	return &FullOrder{db: db, cache: cache}
}

func (f *FullOrder) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	err := f.initDelivery(order.Delivery, order.OrderUID)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}
	err = f.initPayment(order.Payment, order.OrderUID)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}
	order.Items, err = f.initItems(order.Items, order.OrderUID)

	resOrder, err := f.db.CreateOrder(ctx, order)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}
	resDelivery, err := f.db.CreateDelivery(ctx, *order.Delivery)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}

	resPayment, err := f.db.CreatePayment(ctx, *order.Payment)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}
	tmp := make([]model.Item, len(order.Items))
	for i, item := range order.Items {
		res, err := f.db.CreateItem(ctx, item)
		if err != nil {
			log.Println(err)
			return model.Order{}, err
		}
		tmp[i] = res
	}

	return model.Order{
		OrderUID:          resOrder.OrderUID,
		TrackNumber:       resOrder.TrackNumber,
		Entry:             resOrder.Entry,
		Locale:            resOrder.Locale,
		Delivery:          &resDelivery,
		Payment:           &resPayment,
		Items:             tmp,
		InternalSignature: resOrder.InternalSignature,
		CustomerID:        resOrder.CustomerID,
		DeliveryService:   resOrder.DeliveryService,
		ShardKey:          resOrder.ShardKey,
		SMID:              resOrder.SMID,
		DateCreated:       resOrder.DateCreated.UTC(),
		OOFShard:          resOrder.OOFShard,
	}, nil
}

func (f *FullOrder) initDelivery(delivery *model.Delivery, orderUID string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	delivery.ID = id.String()
	delivery.OrderUID = orderUID
	return nil
}

func (f *FullOrder) initPayment(payment *model.Payment, orderUID string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	payment.OrderUID = orderUID
	payment.ID = id.String()
	return nil
}

func (f *FullOrder) initItems(items []model.Item, orderUID string) ([]model.Item, error) {
	tmp := make([]model.Item, len(items))

	for i, item := range items {
		id, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		item.ID = id.String()
		item.OrderUID = orderUID
		tmp[i] = item
	}
	return tmp, nil
}

func (f *FullOrder) GetOrder(ctx context.Context, id string) (model.Order, error) {

	res, err := f.cache.GetOrderByOrderUID(ctx, id)
	if err == nil {
		log.Println("uc redis:", res)
		return res, nil
	}

	resOrder, err := f.db.GetOrderByOrderUID(ctx, id)
	if err != nil {
		return model.Order{}, err
	}

	resDelivery, err := f.db.GetDeliveryByOrderUID(ctx, id)
	if err != nil {
		return model.Order{}, err
	}

	resPayment, err := f.db.GetPaymentsByOrderUID(ctx, id)

	if err != nil {
		return model.Order{}, err
	}

	items, err := f.db.GetItemsByOrderUID(ctx, id)
	if err != nil {
		return model.Order{}, err
	}
	log.Println("uc:", resDelivery)
	log.Println("uc:", resPayment)
	log.Println("uc:", items)
	result := model.Order{
		OrderUID:          resOrder.OrderUID,
		TrackNumber:       resOrder.TrackNumber,
		Entry:             resOrder.Entry,
		Locale:            resOrder.Locale,
		Delivery:          &resDelivery,
		Payment:           &resPayment,
		Items:             items,
		InternalSignature: resOrder.InternalSignature,
		CustomerID:        resOrder.CustomerID,
		DeliveryService:   resOrder.DeliveryService,
		ShardKey:          resOrder.ShardKey,
		SMID:              resOrder.SMID,
		DateCreated:       resOrder.DateCreated,
		OOFShard:          resOrder.OOFShard,
	}
	_, _ = f.cache.CreateOrder(ctx, result)
	log.Println("uc:", result)
	return result, nil
}

func (f *FullOrder) PreloadCache(ctx context.Context) error {
	orders, err := f.db.GetAllOrders(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, o := range orders {
		orderUID := o.OrderUID

		delivery, err := f.db.GetDeliveryByOrderUID(ctx, orderUID)
		if err != nil {
			log.Println("failed to load delivery:", err)
			continue
		}

		payment, err := f.db.GetPaymentsByOrderUID(ctx, orderUID)
		if err != nil {
			log.Println("failed to load payment:", err)
			continue
		}

		items, err := f.db.GetItemsByOrderUID(ctx, orderUID)
		if err != nil {
			log.Println("failed to load items:", err)
			continue
		}

		o.Delivery = &delivery
		o.Payment = &payment
		o.Items = items

		if _, err := f.cache.CreateOrder(ctx, o); err != nil {
			log.Println("failed to cache order:", err)
		}
	}

	return nil
}
