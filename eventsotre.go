package eventstore

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

var (
	ErrDurableNameNotSet = fmt.Errorf("durable name not set")
	ErrHandlerNotSet     = fmt.Errorf("handler not set")
	ErrMessageNotSet     = fmt.Errorf("message not set")
	ErrSubjectNotSet     = fmt.Errorf("subject not set")
)

// Subscription is a basic struct to define
// single subscription
type Subscription struct {
	Sequence    uint64 // optional, default is zero (beginning of data events)
	Subject     string
	DurableName string
	QueueName   string // optional
	Handler     Handler
	Message     proto.Message
}

// Handler is a function which will be called once the the subject type is received
type Handler func(message proto.Message, sequenceID uint64, correlationID, signature string)

// Unsubscribe removes the handler from EventStore
// it should clean up durable name az well
type Unsubscribe func() error

// EventStore base interface
type EventStore interface {
	Publish(message proto.Message, subject, correlationID, signature fmt.Stringer) error
	Subscribe(subOptions *Subscription) (Unsubscribe, error)
	Close() error
}
