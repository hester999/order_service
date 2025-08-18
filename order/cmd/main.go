package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"order/internal/apperr"
	"order/internal/consumer"
	kf "order/internal/consumer/kafka"
	"order/internal/db"
	"order/internal/handler/order"
	"order/internal/repo/inmemmory/redis"
	"order/internal/repo/postgress"
	uc "order/internal/usecases/order"
	"sync"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()

	conn, err := db.Connection()
	if err != nil {
		panic(err)
	}

	client, err := redis.NewRedisClient("redis:6379", ctx)
	if err != nil {
		panic(err)
	}

	cfg := kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "orders",
		GroupID: "order",
	}

	cache := redis.NewRedisRepo(client)
	storage := postgress.NewPostgresStorage(conn)

	usecase := uc.NewFullOrder(storage, cache)
	handler := order.NewOrderHandler(usecase)

	err = usecase.PreloadCache(ctx)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
		} else {
			log.Fatal(err)
		}
	}

	reader := kf.NewKafkaReader(kafka.NewReader(cfg))
	consumer := consumer.NewConsumer(reader, usecase)

	router := mux.NewRouter()
	router.HandleFunc("/orders/{id}", handler.GetOrder).Methods("GET")

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := consumer.Start(ctx); err != nil {
			panic(err)
		}
	}()

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

	wg.Wait()
}
