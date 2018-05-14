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

type Subject interface {
	Validate(string) bool
	String() string
}

// Subscription is a basic struct to define
// single subscription
type Subscription struct {
	Subject     Subject
	DurableName string
	Handler     Handler
	Message     proto.Message
	Timeout     time.Duration // optional, default will be any default timeout realted to implementation
	Sequence    uint64        // optional, default is zero (beginning of data events)
	QueueName   string        // optional
}

type Payload struct {
	Message       proto.Message
	SequenceID    uint64
	CorrelationID string
	Signature     string
	Timestamp     int64 // this value will be set when a message received. It has a noop on Publish
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
	Publish(payload *Payload, subject Subject) error
	Subscribe(subOptions *Subscription) (Unsubscribe, error)
	Close() error
}
