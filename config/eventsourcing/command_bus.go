package eventsourcing

import "github.com/looplab/eventhorizon/commandhandler/bus"

func SetupCommandBus() *bus.CommandHandler {
	return bus.NewCommandHandler()
}
