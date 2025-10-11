package dto

import (
	"order/internal/model"
	"time"
)

func ToModelOrder(order OrderDTO) model.Order {
	m := model.Order{
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

	// --- Delivery ---
	if order.Delivery != nil {
		m.Delivery = &model.Delivery{
			Name:    order.Delivery.Name,
			Phone:   order.Delivery.Phone,
			ZIP:     order.Delivery.ZIP,
			City:    order.Delivery.City,
			Address: order.Delivery.Address,
			Region:  order.Delivery.Region,
			Email:   order.Delivery.Email,
		}
	}

	// --- Payment ---
	if order.Payment != nil {
		m.Payment = &model.Payment{
			Transaction:  order.Payment.Transaction,
			RequestID:    order.Payment.RequestID,
			Currency:     order.Payment.Currency,
			Provider:     order.Payment.Provider,
			Amount:       order.Payment.Amount,
			PaymentDT:    time.Unix(order.Payment.PaymentDT, 0).UTC(),
			Bank:         order.Payment.Bank,
			DeliveryCost: order.Payment.DeliveryCost,
			GoodsTotal:   order.Payment.GoodsTotal,
			CustomFee:    order.Payment.CustomFee,
		}
	}

	// --- Items ---
	m.Items = make([]model.Item, 0, len(order.Items))
	for _, i := range order.Items {
		m.Items = append(m.Items, model.Item{
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

	return m
}
