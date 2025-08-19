package order

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"order/internal/apperr"
	"order/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type OrderRepository struct {
	client *redis.Client
}

func NewOrderRepository(client *redis.Client) *OrderRepository {
	return &OrderRepository{
		client: client,
	}
}

func (o *OrderRepository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	dto := OrderDTO{
		OrderUID:          order.OrderUID,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		ShardKey:          order.ShardKey,
		SMID:              order.SMID,
		DateCreated:       order.DateCreated,
		OOFShard:          order.OOFShard,
	}

	if order.Delivery != nil {
		dto.Delivery = &DeliveryDTO{
			ID:       order.Delivery.ID,
			OrderUID: order.OrderUID,
			Name:     order.Delivery.Name,
			Phone:    order.Delivery.Phone,
			ZIP:      order.Delivery.ZIP,
			City:     order.Delivery.City,
			Address:  order.Delivery.Address,
			Region:   order.Delivery.Region,
			Email:    order.Delivery.Email,
		}
	}

	if order.Payment != nil {
		dto.Payment = &PaymentDTO{
			ID:           order.Payment.ID,
			OrderUID:     order.OrderUID,
			Transaction:  order.Payment.Transaction,
			RequestID:    order.Payment.RequestID,
			Currency:     order.Payment.Currency,
			Provider:     order.Payment.Provider,
			Amount:       order.Payment.Amount,
			PaymentDT:    order.Payment.PaymentDT.Unix(),
			Bank:         order.Payment.Bank,
			DeliveryCost: order.Payment.DeliveryCost,
			GoodsTotal:   order.Payment.GoodsTotal,
			CustomFee:    order.Payment.CustomFee,
		}
	}

	for _, item := range order.Items {
		dto.Items = append(dto.Items, ItemDTO{
			ID:          item.ID,
			OrderUID:    order.OrderUID,
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			RID:         item.RID,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		})
	}

	data, err := json.Marshal(dto)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}

	key := "order:" + dto.OrderUID
	ok, err := o.client.SetNX(ctx, key, data, 2*time.Hour).Result()
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}

	if !ok {
		data, err = o.client.Get(ctx, key).Bytes()
		if err != nil {
			log.Println(err)
			return model.Order{}, err
		}
	}

	return order, nil
}

func (o *OrderRepository) GetOrderByOrderUID(ctx context.Context, id string) (model.Order, error) {
	key := "order:" + id
	data, err := o.client.Get(ctx, key).Bytes()
	if err != nil {
		log.Println(err)
		if errors.Is(err, redis.Nil) {
			return model.Order{}, apperr.ErrNotFound
		}
		return model.Order{}, err
	}

	var dto OrderDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		log.Println(err)
		return model.Order{}, err
	}

	var items []model.Item
	for _, item := range dto.Items {
		items = append(items, model.Item{
			ID:          item.ID,
			OrderUID:    dto.OrderUID,
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			RID:         item.RID,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		})
	}

	return model.Order{
		OrderUID:    dto.OrderUID,
		TrackNumber: dto.TrackNumber,
		Entry:       dto.Entry,
		Delivery: &model.Delivery{
			ID:       dto.Delivery.ID,
			OrderUID: dto.OrderUID,
			Name:     dto.Delivery.Name,
			Phone:    dto.Delivery.Phone,
			ZIP:      dto.Delivery.ZIP,
			City:     dto.Delivery.City,
			Address:  dto.Delivery.Address,
			Region:   dto.Delivery.Region,
			Email:    dto.Delivery.Email,
		},
		Payment: &model.Payment{
			ID:           dto.Payment.ID,
			OrderUID:     dto.OrderUID,
			Transaction:  dto.Payment.Transaction,
			RequestID:    dto.Payment.RequestID,
			Currency:     dto.Payment.Currency,
			Provider:     dto.Payment.Provider,
			Amount:       dto.Payment.Amount,
			PaymentDT:    time.Unix(dto.Payment.PaymentDT, 0),
			Bank:         dto.Payment.Bank,
			DeliveryCost: dto.Payment.DeliveryCost,
			GoodsTotal:   dto.Payment.GoodsTotal,
			CustomFee:    dto.Payment.CustomFee,
		},
		Items:             items,
		Locale:            dto.Locale,
		InternalSignature: dto.InternalSignature,
		CustomerID:        dto.CustomerID,
		DeliveryService:   dto.DeliveryService,
		ShardKey:          dto.ShardKey,
		SMID:              dto.SMID,
		DateCreated:       dto.DateCreated,
		OOFShard:          dto.OOFShard,
	}, nil
}
