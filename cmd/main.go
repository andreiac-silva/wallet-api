package main

import (
	"context"

	"go.uber.org/zap"

	// Automatic load environment variables from .env.
	_ "github.com/joho/godotenv/autoload"

	api "wallet-api/api/wallet"
	"wallet-api/config/db"
	"wallet-api/config/eventsourcing"
	"wallet-api/config/log"
	"wallet-api/config/server"
	"wallet-api/domain/usecases"
)

func main() {
	ctx := context.Background()

	log.Setup()

	mongoClient := db.SetupMongoClient()

	outbox := eventsourcing.SetupOutbox(mongoClient)
	eventBus := eventsourcing.SetupEventBus()
	eventStore := eventsourcing.SetupEventStore(outbox)
	commandBus := eventsourcing.SetupCommandBus()
	balanceRepo := eventsourcing.SetupBalanceRepository(mongoClient)
	balanceUc := usecases.NewBalanceUseCase(balanceRepo)
	eventsourcing.Setup(ctx, eventBus, eventStore, outbox, commandBus, balanceRepo, balanceUc)

	handler := api.NewWalletHandler(commandBus)

	srv := server.SetupHTTPServer(handler)
	if err := srv.ListenAndServe(); err != nil {
		zap.S().Error("something went wrong starting http server", "error", err)
	}

	defer eventsourcing.Close(eventBus, eventStore, outbox, balanceRepo)
}
