package handler

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/delaram-gholampoor-sagha/sd-studio/model"
	"github.com/delaram-gholampoor-sagha/sd-studio/protocol"
)

type OrderHandler struct {
	Service protocol.OrderService
}

func (h *OrderHandler) HandleOrder(msg *message.Message) error {
	var order model.Order
	err := json.Unmarshal(msg.Payload, &order)
	if err != nil {
		return err
	}

	return h.Service.CreateOrder(order)
}
