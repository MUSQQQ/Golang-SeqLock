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
	Counter uint32
	sync.Mutex
}

// reads counter using atomic operations
// used before accessing data protected by the lock
func (seq *SeqLock) RdRead() uint32 {
	return atomic.LoadUint32(&seq.Counter)
}

// checks if data is not being modified by writer
// or if it was not modified since rdRead func
func (seq *SeqLock) RdAgain(val uint32) bool {
	return (atomic.LoadUint32(&seq.Counter)&1) != 0 || val != seq.Counter
}

// resets counter to zero
func (seq *SeqLock) ResetCounter() {
	seq.Lock()
	atomic.SwapUint32(&seq.Counter, 0)
	seq.Lock()
}

// counter becomes odd when writer
// starts modifying data
func (seq *SeqLock) WrLock() {
	seq.Lock()
	atomic.AddUint32(&seq.Counter, 1)
}

// counter becomes even when writer
// starts modifying data
func (seq *SeqLock) WrUnlock() {
	atomic.AddUint32(&seq.Counter, 1)
	seq.Unlock()
}

// locks the data for both writers and readers
// for the given amount of miliseconds
func (seq *SeqLock) TimedBlocker(ms uint64) {
	seq.Lock()
	atomic.AddUint32(&seq.Counter, 1)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	atomic.AddUint32(&seq.Counter, 1)
	seq.Unlock()
}
