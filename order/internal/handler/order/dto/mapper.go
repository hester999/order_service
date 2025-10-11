package dto

import "order/internal/model"

func FromModel(order model.Order) OrderDTO {
	return OrderDTO{
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
		Delivery:          deliveryToDto(order.Delivery),
		Payment:           paymentToDto(order.Payment),
		Items:             itemsToDto(order.Items),
	}
}

func deliveryToDto(d *model.Delivery) *DeliveryDTO {
	if d == nil {
		return nil
	}
	return &DeliveryDTO{
		Name:    d.Name,
		Phone:   d.Phone,
		ZIP:     d.ZIP,
		City:    d.City,
		Address: d.Address,
		Region:  d.Region,
		Email:   d.Email,
	}
}

func paymentToDto(p *model.Payment) *PaymentDTO {
	if p == nil {
		return nil
	}
	return &PaymentDTO{
		Transaction:  p.Transaction,
		RequestID:    p.RequestID,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDT:    p.PaymentDT.Unix(),
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}

func itemsToDto(items []model.Item) []ItemDTO {
	res := make([]ItemDTO, 0, len(items))
	for _, i := range items {
		res = append(res, ItemDTO{
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
	return res
}
