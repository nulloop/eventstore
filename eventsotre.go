package eventstore

import (
	"fmt"
	"time"

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
	Timeout     time.Duration // optional, default will be any default timeout realted to implementation
	Sequence    uint64        // optional, default is zero (beginning of data events)
	Subject     string
	DurableName string
	QueueName   string // optional
	Handler     Handler
	Message     proto.Message
}

type Payload struct {
	Message       proto.Message
	SequenceID    uint64
	CorrelationID string
	Signature     string
	Timestamp     int64
}

// Handler is a function which will be called once the the subject type is received
// if an error return, the message will be redeliver again to the same handler or
// another handler
type Handler func(payload *Payload) error

// Unsubscribe removes the handler from EventStore
// it should clean up durable name az well
type Unsubscribe func() error

// EventStore base interface
type EventStore interface {
	Publish(message proto.Message, subject, correlationID, signature string) error
	Subscribe(subOptions *Subscription) (Unsubscribe, error)
	Close() error
}
