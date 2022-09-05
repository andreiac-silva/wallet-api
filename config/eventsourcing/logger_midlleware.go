package eventsourcing

import (
	"context"

	"github.com/looplab/eventhorizon"
	"go.uber.org/zap"
)

const logger = "logger"

// LoggingMiddleware is a tiny command handle middleware for logging.
func LoggingMiddleware(h eventhorizon.CommandHandler) eventhorizon.CommandHandler {
	return eventhorizon.CommandHandlerFunc(func(ctx context.Context, cmd eventhorizon.Command) error {
		zap.S().Debugf("command: %v", cmd)
		return h.HandleCommand(ctx, cmd)
	})
}

// Logger is a simple event handler for logging all events.
type Logger struct{}

// HandlerType implements the HandlerType method of the eventhorizon.EventHandler interface.
func (l *Logger) HandlerType() eventhorizon.EventHandlerType {
	return logger
}

// HandleEvent implements the HandleEvent method of the EventHandler interface.
func (l *Logger) HandleEvent(_ context.Context, event eventhorizon.Event) error {
	zap.S().Debugf("event: %v", event)
	return nil
}
