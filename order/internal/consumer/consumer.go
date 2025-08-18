package consumer

import (
	"context"
	"encoding/json"
	"log"
	"order/internal/model"
	"time"
)

type Consumer struct {
	reader   MessageReader
	usecases Order
}

func NewConsumer(reader MessageReader, usecases Order) *Consumer {
	return &Consumer{
		reader:   reader,
		usecases: usecases,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	for {
		data, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var dto OrderDTO
		if err := json.Unmarshal(data, &dto); err != nil {
			log.Println("Error unmarshalling:", err)
			continue
		}

		order := model.Order{
			OrderUID:          dto.OrderUID,
			TrackNumber:       dto.TrackNumber,
			Entry:             dto.Entry,
			Locale:            dto.Locale,
			InternalSignature: dto.InternalSignature,
			CustomerID:        dto.CustomerID,
			DeliveryService:   dto.DeliveryService,
			ShardKey:          dto.ShardKey,
			SMID:              dto.SMID,
			DateCreated:       dto.DateCreated,
			OOFShard:          dto.OOFShard,
			Delivery: &model.Delivery{
				Name:    dto.Delivery.Name,
				Phone:   dto.Delivery.Phone,
				ZIP:     dto.Delivery.ZIP,
				City:    dto.Delivery.City,
				Address: dto.Delivery.Address,
				Region:  dto.Delivery.Region,
				Email:   dto.Delivery.Email,
			},
			Payment: &model.Payment{
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
			Items: make([]model.Item, 0, len(dto.Items)),
		}

		for _, i := range dto.Items {
			order.Items = append(order.Items, model.Item{
				ChrtID:      i.ChrtID,
				TrackNumber: i.TrackNumber,
				Price:       i.Price,
				RID:         i.RID,
				Name:        i.Name,
				Sale:        i.Sale,
				Size:        i.Size,
				TotalPrice:  i.TotalPrice,
				NmID:        i.NmID,
				Brand:       i.Brand,
				Status:      i.Status,
			})
		}

		if _, err := c.usecases.CreateOrder(ctx, order); err != nil {
			log.Println("Error creating order:", err)
			continue
		}
	}
}
