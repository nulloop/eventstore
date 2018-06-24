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

type Payload struct {
	id        string
	subject   Subject
	sequence  uint64
	timestamp int64
	message   proto.Message
}

func (p *Payload) ID() string {
	return p.id
}

func (p *Payload) Subject() Subject {
	return p.subject
}

func (p *Payload) Sequence() uint64 {
	return p.sequence
}

func (p *Payload) Timestamp() int64 {
	return p.timestamp
}

func (p *Payload) Message() proto.Message {
	return p.message
}

func NewPayload(subject Subject, message proto.Message, id string) Container {
	payload := &Payload{
		id:      id,
		subject: subject,
		message: message,
	}

	return payload
}
