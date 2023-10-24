package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/delaram-gholampoor-sagha/sd-studio/model"
	"github.com/delaram-gholampoor-sagha/sd-studio/storage"
)

type Order struct {
	redis     *storage.RedisStorage
	publisher message.Publisher
}

func NewOrder(redis *storage.RedisStorage, publisher message.Publisher) *Order {
	return &Order{
		redis:     redis,
		publisher: publisher,
	}
}
func (o *Order) CreateOrder(order model.Order) error {
	ctx := context.Background()

	// Check if order already exists
	existing, err := o.redis.GetCounter(ctx, order.ID)
	if err == nil && existing > 0 {
		return errors.New("order already exists")
	}

	// Store order in Redis (as JSON)
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := o.redis.Set(ctx, order.ID, string(orderJSON), 0); err != nil {
		return fmt.Errorf("failed to save order: %v", err)
	}

	// Publish order-created event with retry mechanism
	msg := message.NewMessage(watermill.NewUUID(), orderJSON)
	if err := o.tryPublish("order-created", msg); err != nil {
		// If after retries the publish still fails, rollback the order save in Redis
		rollbackErr := o.redis.Delete(ctx, order.ID)
		if rollbackErr != nil {
			return fmt.Errorf("failed to rollback order after publish failure: %v, original error: %v", rollbackErr, err)
		}
		return err
	}

	return nil
}

const maxRetries = 3
const retryDelay = 5 * time.Second

// tryPublish attempts to publish a message with retries
func (o *Order) tryPublish(topic string, msg *message.Message) error {
	var lastError error
	for i := 0; i < maxRetries; i++ {
		if err := o.publisher.Publish(topic, msg); err == nil {
			return nil
		} else {
			lastError = err
			time.Sleep(retryDelay)
		}
	}
	return fmt.Errorf("failed to publish message after %d attempts: %v", maxRetries, lastError)
}
