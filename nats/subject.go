package nats

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
)

var (
	ErrTopicEmpty = fmt.Errorf("topic is empty")
)

type Subject struct {
	topic       string
	durable     string
	queue       string
	sequence    uint64
	msgInstance proto.Message
}

func (n *Subject) Topic() string {
	return n.topic
}

type Option func(*Subject)

func OptQueueName(name string) Option {
	return func(subject *Subject) {
		subject.queue = name
	}
}

func OptDurableName(name string) Option {
	return func(subject *Subject) {
		subject.durable = name
	}
}

func OptMessageInstance(msg proto.Message) Option {
	return func(subject *Subject) {
		subject.msgInstance = msg
	}
}

func OptSequence(sequence uint64) Option {
	return func(subject *Subject) {
		subject.sequence = sequence
	}
}

func NewSubject(topic string, options ...Option) (*Subject, error) {
	if topic == "" {
		return nil, ErrTopicEmpty
	}

	subject := &Subject{
		topic: topic,
	}

	for _, option := range options {
		option(subject)
	}

	return subject, nil
}
