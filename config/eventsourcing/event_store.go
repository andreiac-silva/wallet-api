package eventsourcing

import (
	"os"

	"github.com/looplab/eventhorizon"
	esMongo "github.com/looplab/eventhorizon/eventstore/mongodb_v2"
	mongoOutbox "github.com/looplab/eventhorizon/outbox/mongodb"
	"go.uber.org/zap"
)

func SetupEventStore(outbox *mongoOutbox.Outbox) eventhorizon.EventStore {
	dbPrefix := os.Getenv("MONGO_DB_NAME")

	eventStore, err := esMongo.NewEventStoreWithClient(
		outbox.Client(),
		dbPrefix,
		esMongo.WithEventHandlerInTX(outbox),
	)
	if err != nil {
		zap.S().Fatal("event store setup has been failed", "error", err)
	}

	zap.S().Debug("event store setup has been done")

	return eventStore
}
