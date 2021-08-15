package seqlock

import (
	"sync"
	"sync/atomic"
)

// Counter increases by 1 everytime the lock gets taken or released by writer
// sequence number is other name for it
// inherited mutex is used only by writers
type SeqLock struct {
	Counter int32
	sync.Mutex
}

// func that reads counter using atomic operations
// used before accessing data protected by the lock
func (seq *SeqLock) RdRead() int32 {
	return atomic.LoadInt32(&seq.Counter)
}

// func that checks if data is not being modified by writer
// or if it was not modified since rdRead func
func (seq *SeqLock) RdAgain(val int32) bool {
	return (atomic.LoadInt32(&seq.Counter)&1) != 0 || val != seq.Counter
}

// writer lock func, Counter becomes odd when writer
// starts modifying data
func (seq *SeqLock) WrLock() {
	seq.Lock()
	atomic.AddInt32(&seq.Counter, 1)
}

// writer unlock func, Counter becomes even when writer
// starts modifying data
func (seq *SeqLock) WrUnlock() {
	atomic.AddInt32(&seq.Counter, 1)
	seq.Unlock()
}
