package eventstore

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	pb "github.com/alinz/eventstore/proto"
	proto "github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
)

type natsEventStore struct {
	conn stan.Conn
}

func (n *natsEventStore) Publish(message proto.Message, subject, coorlationID, signature fmt.Stringer) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	data, err := proto.Marshal(&pb.Transport{
		CoorlationId: coorlationID.String(),
		Signature:    signature.String(),
		Payload:      payload,
	})
	if err != nil {
		return err
	}

	return n.conn.Publish(subject.String(), data)
}

func (n *natsEventStore) Subscribe(subscription *Subscription) (Unsubscribe, error) {
	if subscription.DurableName == "" {
		return nil, fmt.Errorf("durable name not set")
	}

	if subscription.Handler == nil {
		return nil, fmt.Errorf("handler not set")
	}

	if subscription.Message == nil {
		return nil, fmt.Errorf("message not set")
	}

	if subscription.Subject == "" {
		return nil, fmt.Errorf("subject not set")
	}

	transport := pb.Transport{}
	process := func(msg *stan.Msg) {
		err := proto.Unmarshal(msg.Data, &transport)
		if err != nil {
			log.Printf("error decoding transport, %v, %s\n", msg, err)
			return
		}

		err = proto.Unmarshal(transport.Payload, subscription.Message)
		if err != nil {
			log.Printf("error decoding message, %v, %s\n", transport, err)
			return
		}

		subscription.Handler(subscription.Message, msg.Sequence, transport.CoorlationId, transport.Signature)
	}

	var err error
	var subscriptionHandler stan.Subscription

	var options []stan.SubscriptionOption
	options = append(options, stan.DurableName(subscription.DurableName))
	if subscription.Sequence != 0 {
		options = append(options, stan.StartAtSequence(subscription.Sequence))
	}

	if subscription.QueueName == "" {
		subscriptionHandler, err = n.conn.Subscribe(subscription.Subject, process, options...)
	} else {
		subscriptionHandler, err = n.conn.QueueSubscribe(subscription.Subject, subscription.QueueName, process, options...)
	}

	if err != nil {
		return nil, err
	}

	return func() error { return subscriptionHandler.Close() }, nil
}

func (n *natsEventStore) Close() error {
	return n.conn.Close()
}

// NewNatsStreaming creates a nee eventstore
func NewNatsStreaming(tlsConfig *tls.Config, addr, clusterID, clientID string) (EventStore, error) {
	nc, err := nats.Connect(addr, nats.Secure(tlsConfig))
	if err != nil {
		return nil, err
	}

	var conn stan.Conn

	for {
		conn, err = stan.Connect(clusterID, clientID, stan.NatsConn(nc))
		if err != nil {
			if err == stan.ErrConnectReqTimeout {
				log.Printf("trying to connect to nats-streaming at %s\n", addr)
				time.Sleep(1 * time.Second)
				continue
			}
			return nil, err
		}

		break
	}

	return &natsEventStore{
		conn,
	}, nil
}
