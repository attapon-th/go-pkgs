package task

import (
	"sync"
	"sync/atomic"
)

// WaitGroupCount like sync.WaitGroup, but overwrite Add(), Done(), GetCount()
type WaitGroupCount struct {
	sync.WaitGroup
	count int64
}

// Add - Add add one, which may be negative, to the WaitGroupCount counter.
func (wg *WaitGroupCount) Add() {
	atomic.AddInt64(&wg.count, 1)
	wg.WaitGroup.Add(1)
}

// Done - decrements the WaitGroupCount counter by one
func (wg *WaitGroupCount) Done() {
	atomic.AddInt64(&wg.count, -1)
	wg.WaitGroup.Done()
}

// GetCount - Get WaitGroupCount counter
func (wg *WaitGroupCount) GetCount() int {
	return int(atomic.LoadInt64(&wg.count))
}

// DoneAll - done process all **force kill**
func (wg *WaitGroupCount) DoneAll() {
	for {
		if wg.GetCount() <= 0 {
			break
		}
		wg.Done()
	}
}
