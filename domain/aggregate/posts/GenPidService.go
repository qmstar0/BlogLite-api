package posts

import (
	"sync"
	"sync/atomic"
)

var GenPId = newAtomicIDGenerator(10000)

type AtomicIDGenerator struct {
	value int64
	mu    sync.Mutex
}

func newAtomicIDGenerator(initialValue int64) *AtomicIDGenerator {
	return &AtomicIDGenerator{value: initialValue}
}

func (g *AtomicIDGenerator) NextID() int {
	return int(atomic.AddInt64(&g.value, 1))
}
