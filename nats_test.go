package eventstore_test

import (
	"math/rand"
	"sync"
	"testing"

	proto "github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/nats-io/nats-streaming-server/server"
	"github.com/nulloop/eventstore"
	pb "github.com/nulloop/eventstore/proto"
)

type genRandomStringer struct{}

func (genRandomStringer) String() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var randomStringer genRandomStringer

func runDummyServer(clusterName string) (*server.StanServer, error) {
	s, err := server.RunServer(clusterName)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func TestRunDummyServer(t *testing.T) {
	server, err := runDummyServer("dummy_server")
	if err != nil {
		t.Error(err)
	}
	server.Shutdown()
}

func TestCreateEventStore(t *testing.T) {
	server, err := runDummyServer("dummy")
	if err != nil {
		t.Error(err)
	}

	defer server.Shutdown()

	es, err := eventstore.NewNatsStreaming(nil, nats.DefaultURL, "dummy", "client1")
	if err != nil {
		t.Error(err)
	}

	defer es.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	unsubscribe, err := es.Subscribe(&eventstore.Subscription{
		Message:     &pb.DummyMessage{},
		Subject:     pb.Dummy_Subject1.String(),
		DurableName: pb.Dummy_DurableName1.String(),
		QueueName:   pb.Dummy_QueueName1.String(),
		Handler: func(message proto.Message, sequenceID uint64, correlationID, signature string) {
			dummyMessage, ok := message.(*pb.DummyMessage)
			if !ok {
				t.Error("message is not DummyMessage")
			}

			if dummyMessage.Value != "this is test" {
				t.Error("message is incorrect")
			}

			wg.Done()
		},
	})

	if err != nil {
		t.Error(err)
	}

	defer unsubscribe()

	es.Publish(&pb.DummyMessage{Value: "this is test"}, pb.Dummy_Subject1, randomStringer, randomStringer)

	wg.Wait()
}
