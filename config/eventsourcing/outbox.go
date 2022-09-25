package eventsourcing

import (
	"os"

	mongoOutbox "github.com/looplab/eventhorizon/outbox/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SetupOutbox(client *mongo.Client) *mongoOutbox.Outbox {
	dbName := os.Getenv("MONGO_DB_NAME")
	outbox, err := mongoOutbox.NewOutboxWithClient(client, dbName)
	if err != nil {
		zap.S().Fatal("outbox setup has been failed", "error", err)
	}

	go func() {
		for e := range outbox.Errors() {
			zap.S().Errorw("there are errors on outbox flow", "error", e)
		}
	}()

	zap.S().Debug("outbox setup has been done")

	return outbox
}
