package seqlock

import (
	"sync"
	"sync/atomic"
	"time"
)

// Counter increases by 1 everytime the lock gets taken or released by writer
// sequence number is other name for it
// inherited mutex is used only by writers
type SeqLock struct {
	Counter uint64
	sync.Mutex
}

func NewSeqLock() *SeqLock {
	seqlock := SeqLock{Counter: 0}
	return &seqlock
}

// reads counter using atomic operations
// used before accessing data protected by the lock
func (seq *SeqLock) RdRead() uint64 {
	return atomic.LoadUint64(&seq.Counter)
}

// checks if data is not being modified by writer
// or if it has not been modified since rdRead func
func (seq *SeqLock) RdAgain(val uint64) bool {
	return (atomic.LoadUint64(&seq.Counter)&1) != 0 || val != seq.Counter
}

// resets counter to zero
func (seq *SeqLock) ResetCounter() {
	seq.Lock()
	atomic.SwapUint64(&seq.Counter, 0)
	seq.Unlock()
}

// counter becomes odd when writer
// starts modifying data
func (seq *SeqLock) WrLock() {
	seq.Lock()
	atomic.AddUint64(&seq.Counter, 1)
}

// counter becomes even when writer
// starts modifying data
func (seq *SeqLock) WrUnlock() {
	atomic.AddUint64(&seq.Counter, 1)
	seq.Unlock()
}

// locks the data for both writers and readers
// for the given amount of miliseconds
func (seq *SeqLock) TimeBlock(ms int64) {
	seq.Lock()
	atomic.AddUint64(&seq.Counter, 1)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	atomic.AddUint64(&seq.Counter, 1)
	seq.Unlock()
}
