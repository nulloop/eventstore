package nats

import (
	"fmt"

	"github.com/nulloop/eventstore"
)

var (
	ErrNotNatsEventStore = fmt.Errorf("eventstore not nats one")
)

func PublishEnableFlag(eventstore eventstore.EventStore) error {
	nats, ok := eventstore.(*NatsEventstore)
	if !ok {
		return ErrNotNatsEventStore
	}

	nats.publishEnabled.Set()
	return nil
}

func PublishDisableFlag(eventstore eventstore.EventStore) error {
	nats, ok := eventstore.(*NatsEventstore)
	if !ok {
		return ErrNotNatsEventStore
	}

	nats.publishEnabled.UnSet()
	return nil
}
