package nats

import (
	"fmt"
	"reflect"

	"github.com/gogo/protobuf/proto"
)

var (
	ErrTopicEmpty   = fmt.Errorf("topic is empty")
	ErrMessageEmpty = fmt.Errorf("message is empty")
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

func (n *Subject) Clone(options ...Option) (*Subject, error) {
	subject, err := NewSubject(n.topic)
	if err != nil {
		return nil, err
	}

	subject.durable = ""
	subject.queue = ""
	subject.sequence = 0
	subject.msgInstance = nil

	for _, option := range options {
		option(subject)
	}

	if subject.msgInstance == nil || reflect.ValueOf(subject.msgInstance).IsNil() {
		return nil, ErrMessageEmpty
	}

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
