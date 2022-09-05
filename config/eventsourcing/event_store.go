package eventsourcing

import (
	"os"

	"github.com/looplab/eventhorizon"
	esMongo "github.com/looplab/eventhorizon/eventstore/mongodb_v2"
	"go.uber.org/zap"
)

func SetupEventStore(eventBus eventhorizon.EventBus) eventhorizon.EventStore {
	url := os.Getenv("MONGO_URI")
	dbPrefix := os.Getenv("MONGO_DB_NAME")

	eventStore, err := esMongo.NewEventStore(url, dbPrefix, esMongo.WithEventHandler(eventBus))
	if err != nil {
		zap.S().Fatal("event store setup has been failed", "error", err)
	}

	zap.S().Debug("event store setup has been done")
	return eventStore
}
