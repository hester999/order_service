package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"order/internal/apperr"
	"order/internal/model"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	uc Order
}

func NewOrderHandler(uc Order) *OrderHandler {
	return &OrderHandler{uc}
}

func (o *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("GetOrder")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	res, err := o.uc.GetOrder(ctx, id)
	if err != nil {
		log.Println("GetOrder")
		if errors.Is(err, apperr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
				Code    int    `json:"code"`
			}{Message: "order not found", Code: http.StatusNotFound})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{Message: "internal server error", Code: http.StatusInternalServerError})
		return
	}
	log.Println("handler", res)

	dto := OrderDTO{
		OrderUID:          res.OrderUID,
		TrackNumber:       res.TrackNumber,
		Entry:             res.Entry,
		Locale:            res.Locale,
		InternalSignature: res.InternalSignature,
		CustomerID:        res.CustomerID,
		DeliveryService:   res.DeliveryService,
		ShardKey:          res.ShardKey,
		SMID:              res.SMID,
		DateCreated:       res.DateCreated,
		OOFShard:          res.OOFShard,
		Delivery:          o.deliveryToDto(res.Delivery),
		Payment:           o.paymentToDto(res.Payment),
		Items:             o.itemsToDto(res.Items),
	}

	fmt.Println(dto)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		log.Println("GetOrder")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{Message: "internal server error", Code: http.StatusInternalServerError})
	}
}

func (o *OrderHandler) deliveryToDto(d *model.Delivery) *DeliveryDTO {
	if d == nil {
		return nil
	}
	return &DeliveryDTO{
		ID:       d.ID,
		OrderUID: d.OrderUID,
		Name:     d.Name,
		Phone:    d.Phone,
		ZIP:      d.ZIP,
		City:     d.City,
		Address:  d.Address,
		Region:   d.Region,
		Email:    d.Email,
	}
}

func (o *OrderHandler) paymentToDto(p *model.Payment) *PaymentDTO {
	if p == nil {
		return nil
	}
	return &PaymentDTO{
		ID:           p.ID,
		OrderUID:     p.OrderUID,
		Transaction:  p.Transaction,
		RequestID:    p.RequestID,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDT:    p.PaymentDT.Unix(), // преобразуем в timestamp
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}

func (o *OrderHandler) itemsToDto(items []model.Item) []ItemDTO {
	res := make([]ItemDTO, 0, len(items))
	for _, i := range items {
		res = append(res, ItemDTO{
			ID:          i.ID,
			OrderUID:    i.OrderUID,
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
