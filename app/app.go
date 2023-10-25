package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/delaram-gholampoor-sagha/sd-studio/config"
	"github.com/delaram-gholampoor-sagha/sd-studio/handler"
	"github.com/delaram-gholampoor-sagha/sd-studio/router"
	"github.com/delaram-gholampoor-sagha/sd-studio/services"
	"github.com/delaram-gholampoor-sagha/sd-studio/storage"
	"github.com/delaram-gholampoor-sagha/sd-studio/subscriber"
)

func Run() {
	cfg := config.LoadConfig()

	redisStorage := storage.NewRedisStorage(cfg.RedisAddress, cfg.RedisPassword, cfg.RedisDB)
	counterService := services.NewCounterService(redisStorage)

	logger := watermill.NewStdLogger(false, false)
	publisher := gochannel.NewGoChannel(gochannel.Config{}, logger)

	orderService := services.NewOrder(redisStorage, publisher)

	orderHandler := &handler.OrderHandler{Service: orderService}
	counterSubscriber := &subscriber.CounterSubscriber{CounterService: counterService}

	r, err := router.NewRouter()
	if err != nil {
		log.Fatalf("Could not create router: %v", err)
	}

	subscriber := gochannel.NewGoChannel(gochannel.Config{}, logger)

	go func() {
		messages, err := subscriber.Subscribe(context.Background(), "order-created")
		if err != nil {
			log.Fatalf("Could not subscribe: %v", err)
		}
		for msg := range messages {
			if err := counterSubscriber.HandleOrderCreated(msg, context.Background()); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	r.AddNoPublisherHandler(
		"order-handler",
		"orders",
		subscriber,
		orderHandler.HandleOrder,
	)

	http.HandleFunc("/orders", orderHandler.HTTPCreateHandle)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
	}()

	if err := r.Run(ctx); err != nil {
		log.Fatalf("Could not start router: %v", err)
	}
}
