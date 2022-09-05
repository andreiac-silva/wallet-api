package eventsourcing

import (
	"context"
	"os"

	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventbus/gcp"
	mongoOutbox "github.com/looplab/eventhorizon/outbox/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SetupOutbox(ctx context.Context, client *mongo.Client, eventBus *gcp.EventBus) *mongoOutbox.Outbox {
	dbName := os.Getenv("MONGO_DB_NAME")

	outbox, err := mongoOutbox.NewOutboxWithClient(client, dbName)
	if err != nil {
		zap.S().Fatal("outbox setup has been failed", "error", err)
	}

	go func() {
		for e := range outbox.Errors() {
			zap.S().Error("there are errors on outbox flow", "error", e)
		}
	}()

	if err := outbox.AddHandler(ctx, eventhorizon.MatchAll{}, eventBus); err != nil {
		zap.S().Fatal("failure to add event bus as outbox handler", "error", err)
	}

	zap.S().Debug("outbox setup has been done")
	return outbox
}
