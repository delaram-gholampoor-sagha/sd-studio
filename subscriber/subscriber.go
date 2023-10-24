package subscriber

import (
	"context"

	"github.com/delaram-gholampoor-sagha/sd-studio/protocol"

	"github.com/ThreeDotsLabs/watermill/message"
)

type CounterSubscriber struct {
	CounterService protocol.CounterService
}

func (s *CounterSubscriber) HandleOrderCreated(msg *message.Message, ctx context.Context) error {
	return s.CounterService.Increment(ctx, msg.UUID)
}
