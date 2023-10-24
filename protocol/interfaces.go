package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/sd-studio/model"
)

type OrderService interface {
	CreateOrder(model model.Order) error
}

type CounterService interface {
	Increment(ctx context.Context, key string) error
}
