package nats

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
)

var (
	ErrTopicNotSet      = fmt.Errorf("topic not set")
	ErrMsgBuilderNotSet = fmt.Errorf("message builder not set")
)

type MsgBuilder func() proto.Message

type Subject struct {
	topic       string
	msgBuilder  MsgBuilder
	msgInstance proto.Message

	durable  string
	queue    string
	sequence uint64
}

func (n *Subject) Topic() string {
	return n.topic
}

func (n *Subject) Instance(options ...Option) (*Subject, error) {
	subject, err := NewSubject(n.topic, n.msgBuilder, options...)
	if err != nil {
		return nil, err
	}

	subject.msgInstance = subject.msgBuilder()
	return subject, nil
}

func (n *Subject) UpdateSequence(sequence uint64) {
	n.sequence = sequence
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

func OptSequence(sequence uint64) Option {
	return func(subject *Subject) {
		subject.sequence = sequence
	}
}

func NewSubject(topic string, msgBuilder MsgBuilder, options ...Option) (*Subject, error) {
	if topic == "" {
		return nil, ErrTopicNotSet
	}

	subject := &Subject{
		topic:      topic,
		msgBuilder: msgBuilder,
	}

	for _, option := range options {
		option(subject)
	}

	return subject, nil
}
