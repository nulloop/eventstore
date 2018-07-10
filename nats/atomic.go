// this code was part of
// https://github.com/tevino/abool

package nats

import (
	"sync/atomic"
)

type AtomicBool int32

func (ab *AtomicBool) Set() {
	atomic.StoreInt32((*int32)(ab), 1)
}

func (ab *AtomicBool) UnSet() {
	atomic.StoreInt32((*int32)(ab), 0)
}

func (ab *AtomicBool) Value() bool {
	return atomic.LoadInt32((*int32)(ab)) == 1
}

func NewAtomicBool() *AtomicBool {
	return new(AtomicBool)
}
