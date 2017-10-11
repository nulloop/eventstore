package eventstore

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

// Subscription is a basic struct to define
// single subscription
type Subscription struct {
	Sequence    uint64
	Subject     string
	DurableName string
	QueueName   string
	Handler     Handler
	Message     proto.Message
}

// Handler is a function which will be called once the the subject type is received
type Handler func(message proto.Message, sequenceID uint64, coorlationID, signature string)

// Unsubscribe removes the handler from EventStore
// it should clean up durable name az well
type Unsubscribe func() error

// EventStore base interface
type EventStore interface {
	Publish(message proto.Message, subject, coorlationID, signature fmt.Stringer) error
	Subscribe(subOptions *Subscription) (Unsubscribe, error)
	Close() error
}
