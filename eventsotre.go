package eventstore

import (
	proto "github.com/golang/protobuf/proto"
)

// Container is a common interface for eventstore event's payload
type Container interface {
	ID() string // ID can be used as correlation
	Subject() Subject
	Sequence() uint64
	Timestamp() int64 // this value will be set when a message received. It has a noop on Publish
	Message() proto.Message
	ActiveMode() bool
}

// Subject is a function type which
type Subject interface {
	Topic() string
}

// Unsubscribe removes the handler from EventStore
// it should clean up durable name az well
type Unsubscribe func() error

type Handler func(Container) error

// EventStore base interface
type EventStore interface {
	Publish(Container) error
	Subscribe(Subject, Handler) (Unsubscribe, error)
	Close() error
}
