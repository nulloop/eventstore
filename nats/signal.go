package nats

import (
	"github.com/nulloop/eventstore"
)

// SignalCond is a function which must be used
// and return true if it needs to make the Push call noop
type SignalCond func(eventstore.Container) bool

// Signal is Special type which helps implement
// the warm up feature in event store.
//
// There are multiple write which calls Push method and one condition logic which
// consume each pushed messages and if it returns true, all the Push method becomes NOOP
type Signal struct {
	done chan struct{}
	data chan eventstore.Container
}

// Push pushes message down to be checked. It becomes NOOP once the
// condition becomes satisfied
func (s *Signal) Push(a eventstore.Container) {
	select {
	case <-s.done:
	case s.data <- a:
	}
}

func (s *Signal) next() (eventstore.Container, bool) {
	select {
	case val, ok := <-s.data:
		return val, ok
	}
}

// NewSignal creats Signal object and accpet cond func
// once the cond returns true, all the Push calls becomes noop
func NewSignal(cond SignalCond, success func()) *Signal {
	signal := &Signal{
		done: make(chan struct{}, 1),
		data: make(chan eventstore.Container, 1),
	}

	go func() {
		defer close(signal.done)

		for {
			val, ok := signal.next()
			if ok {
				if cond(val) {
					success()
					break
				}
			}
		}
	}()

	return signal
}
