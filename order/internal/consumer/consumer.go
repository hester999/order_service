package consumer

import (
	"context"
	"encoding/json"
	"log"
	"order/internal/consumer/dto"
	"order/internal/consumer/validator"
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

		select {
		case <-ctx.Done():
			log.Println("Consumer shutting down gracefully...")
			return ctx.Err()
		default:
		}
		data, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var orderDTO dto.OrderDTO
		if err := json.Unmarshal(data, &orderDTO); err != nil {
			log.Println("Error unmarshalling order:", err)
			continue
		}

		err = validator.ValidateIncomingOrder(orderDTO)
		if err != nil {
			log.Println("Error validating order:", err)
			continue
		}

		order := dto.ToModelOrder(orderDTO)

		log.Printf("Processing order: %s\n", order.OrderUID)
		if _, err := c.usecases.CreateOrder(ctx, order); err != nil {
			log.Printf("Error saving order %s: %v\n", order.OrderUID, err)
			continue
		}

	}
}
