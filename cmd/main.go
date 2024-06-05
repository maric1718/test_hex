package main

import (
	"fmt"
	"pos/internal/adapter/config"
	"pos/internal/adapter/handler/http"
	"pos/internal/adapter/logger"
	"pos/internal/adapter/storage/file_system/repository"
	"pos/internal/adapter/storage/kafka"
	"pos/internal/core/service"

	"log/slog"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		return
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// TEMP
	// // Init database
	// db, err := postgres.New(config.DB)
	// if err != nil {
	// 	slog.Error("Error initializing database connection", "error", err)
	// 	return
	// }

	// slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// // Migrate database
	// err = db.Migrate()
	// if err != nil {
	// 	slog.Error("Error migrating database", "error", err)
	// 	return
	// }

	kafkaMarket := kafka.KafkaConnection(config.Kafka.Timeout, config.Kafka.URL, config.Kafka.Topic)
	kafkaSvc := service.NewKafkaService(kafkaMarket)

	// TEMP
	go func() { // just assume that this is another service that register kafka topic
		kafkaSvc.ReadMessage()
	}()

	// TEMP
	// initEvents, err := domain.LoadInitEvents()
	// if err != nil {
	// 	slog.Error("Error initializing data", "error", err)
	// 	return
	// }

	// out, _ := json.MarshalIndent(initEvents, "", " ")
	// fmt.Println(string(out))

	// initMarkets, err := domain.LoadInitMarkets()
	// if err != nil {
	// 	slog.Error("Error initializing data", "error", err)
	// 	return
	// }

	// fmt.Println(initMarkets)

	marketRepo := repository.NewMarketRepository()
	marketService := service.NewMarketService(marketRepo)
	marketHandler := http.NewMarketHandler(marketService)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		// token,
		*marketHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		return
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		return
	}
}
