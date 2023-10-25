package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/delaram-gholampoor-sagha/sd-studio/model"
	"github.com/delaram-gholampoor-sagha/sd-studio/protocol"
)

type OrderHandler struct {
	Service          protocol.OrderService
	HTTPCreateHandle func(w http.ResponseWriter, r *http.Request)
}

func (h *OrderHandler) HandleOrder(msg *message.Message) error {
	var order model.Order
	err := json.Unmarshal(msg.Payload, &order)
	if err != nil {
		return err
	}

	return h.Service.CreateOrder(order)
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {

	var order model.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateOrder(order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order created successfully"))
}
