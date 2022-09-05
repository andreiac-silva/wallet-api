package eventsourcing

import (
	"os"

	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/repo/mongodb"
	"github.com/looplab/eventhorizon/repo/version"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"wallet-api/domain/projections"
)

type BalanceRepository struct {
	eventhorizon.ReadWriteRepo
}

func SetupBalanceRepository(mongoClient *mongo.Client) BalanceRepository {
	dbName := os.Getenv("MONGO_DB_NAME")
	collection := os.Getenv("MONGO_COLLECTION_BALANCES")

	repo, err := mongodb.NewRepoWithClient(mongoClient, dbName, collection)
	if err != nil {
		zap.S().Fatal("balance repository has been failed", "error", err)
	}

	repo.SetEntityFactory(func() eventhorizon.Entity {
		return &projections.Balance{}
	})

	balanceRepo := version.NewRepo(repo)

	zap.S().Debug("balance repository setup has been done")
	return BalanceRepository{ReadWriteRepo: balanceRepo}
}
