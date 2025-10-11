package order

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"order/internal/apperr"
	"order/internal/handler/order/dto"

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
			err = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
				Code    int    `json:"code"`
			}{Message: "order not found", Code: http.StatusNotFound})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{Message: "internal server error", Code: http.StatusInternalServerError})
		return
	}

	modelResponse := dto.FromModel(res)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(modelResponse)
	if err != nil {
		log.Println("GetOrder")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{Message: "internal server error", Code: http.StatusInternalServerError})
	}
}
