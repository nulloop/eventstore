package eventstore

import (
	"crypto/tls"
	"log"
	"time"

	proto "github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	pb "github.com/nulloop/eventstore/proto"
)

type natsEventStore struct {
	conn stan.Conn
}

func (n *natsEventStore) Publish(message proto.Message, subject, correlation, signature string) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	data, err := proto.Marshal(&pb.Transport{
		CorrelationId: correlation,
		Signature:     signature,
		Payload:       payload,
	})
	if err != nil {
		return err
	}

	return n.conn.Publish(subject, data)
}

func (n *natsEventStore) Subscribe(subscription *Subscription) (Unsubscribe, error) {
	if subscription.DurableName == "" {
		return nil, ErrDurableNameNotSet
	}

	if subscription.Handler == nil {
		return nil, ErrHandlerNotSet
	}

	if subscription.Message == nil {
		return nil, ErrMessageNotSet
	}

	if subscription.Subject == "" {
		return nil, ErrSubjectNotSet
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

		err = subscription.Handler(&Payload{
			Message:       subscription.Message,
			SequenceID:    msg.Sequence,
			CorrelationID: transport.CorrelationId,
			Signature:     transport.Signature,
			Timestamp:     msg.Timestamp,
		})

		if err == nil {
			if err = msg.Ack(); err != nil {
				log.Printf("error ack message for %v, %s\n", msg, err)
			}
		}
	}

	var err error
	var subscriptionHandler stan.Subscription

	options := []stan.SubscriptionOption{
		stan.DurableName(subscription.DurableName),
		stan.StartAtSequence(subscription.Sequence),
		stan.SetManualAckMode(),
	}

	// aw, _ := time.ParseDuration("5s")
	if subscription.Timeout != 0 {
		options = append(options, stan.AckWait(subscription.Timeout))
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
	opts := make([]nats.Option, 0)
	if tlsConfig != nil {
		opts = append(opts, nats.Secure(tlsConfig))
	}

	nc, err := nats.Connect(addr, opts...)
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
