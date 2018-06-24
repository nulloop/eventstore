package nats

import (
	"github.com/golang/protobuf/proto"
	"github.com/nulloop/eventstore"
)

type Payload struct {
	id        string
	subject   eventstore.Subject
	sequence  uint64
	timestamp int64
	message   proto.Message
}

func (p *Payload) ID() string {
	return p.id
}

func (p *Payload) Subject() eventstore.Subject {
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

func NewPayload(subject eventstore.Subject, message proto.Message) *Payload {
	payload := &Payload{
		subject: subject,
		message: message,
	}

	return payload
}
