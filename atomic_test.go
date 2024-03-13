package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

const counterGoroutines = 1_000_000

func initMutexCounter(
	goroutinesCounter int,
) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(goroutinesCounter)
	for i := 0; i < goroutinesCounter; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter += 1
		}()
	}
	wg.Wait()
}

func initAtomicCounter(
	goroutinesCounter int,
) {
	var wg sync.WaitGroup
	wg.Add(goroutinesCounter)
	for i := 0; i < goroutinesCounter; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicCounter(b *testing.B) {
	counter = 0
	initAtomicCounter(counterGoroutines)
}

func BenchmarkMutexCounter(b *testing.B) {
	counter = 0
	initMutexCounter(counterGoroutines)
}
