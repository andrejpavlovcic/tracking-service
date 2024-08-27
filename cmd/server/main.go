package main

import (
	"fmt"
	"net/http"
	"tracking_system/internal/db"
	"tracking_system/internal/kafka"
	"tracking_system/internal/logger"
	"tracking_system/internal/redis"
	"tracking_system/internal/routes"
)

func main() {
	// Init logger
	logger.IntitLogger()

	// Init Kafka Producer
	kafka.InitKafkaProducer()

	// Init DB
	if err := db.InitDB(); err != nil {
		panic(err)
	}

	// Init Redis
	redis.InitRedis()

	// Init Router
	router := routes.NewRouter()

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
