package main

import (
	"sync"
	"sync/atomic"
	"time"
)

// using mutex
func MutexCounter() int {
	// go routines counter
	goroutinesCount := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				goroutinesCount--
				wg.Done()
				m.Unlock()
			}()

			m.Lock()
			goroutinesCount++
			time.Sleep(time.Microsecond)
		}()
	}

	wg.Wait()
	return goroutinesCount
}

// using atomic
func AtomicCounter() int32 { // exact type, not just int
	goroutinesCount := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				atomic.AddInt32(&goroutinesCount, -1)
				wg.Done()
			}()

			atomic.AddInt32(&goroutinesCount, 1)
			time.Sleep(time.Microsecond)
		}()
	}

	wg.Wait()
	return goroutinesCount
}
