package eventsourcing

import (
	"os"

	"github.com/looplab/eventhorizon/eventbus/gcp"
	"go.uber.org/zap"
)

func SetupEventBus() *gcp.EventBus {
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	appID := os.Getenv("APP_ID")

	eventBus, err := gcp.NewEventBus(projectID, appID)
	if err != nil {
		zap.S().Fatal("event bus setup has been failed", "error", err)
	}

	go func() {
		for e := range eventBus.Errors() {
			zap.S().Error("there are errors on event bus flow", "error", e)
		}
	}()

	zap.S().Debug("event bus setup has been done")

	return eventBus
}
