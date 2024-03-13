package main

import (
	"sync"
	"testing"
	"time"
)

var counter = int64(0)

func locker(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	result := int64(0)
	for i := 0; i < 100; i++ {
		l.Lock()
		result += counter
		time.Sleep(time.Microsecond)
		l.Unlock()
	}
}

func initGoroutines(
	readerCount int,
	writerCount int,
	readerLocker sync.Locker,
	writerLocker sync.Locker,
) {
	var wg sync.WaitGroup
	wg.Add(readerCount + writerCount)
	for i := 0; i < writerCount; i++ {
		go locker(&wg, writerLocker)
	}
	for i := 0; i < readerCount; i++ {
		go locker(&wg, readerLocker)
	}
	wg.Wait()
}

func BenchmarkMutex1000Read1Write(b *testing.B) {
	counter = 0
	var mu sync.Mutex
	initGoroutines(1000, 1, &mu, &mu)
}

func BenchmarkRWMutex1000Read1Write(b *testing.B) {
	counter = 0
	var mu sync.RWMutex
	initGoroutines(1000, 1, mu.RLocker(), &mu)
}

func BenchmarkMutex1000Read100Write(b *testing.B) {
	counter = 0
	var mu sync.Mutex
	initGoroutines(1000, 100, &mu, &mu)
}

func BenchmarkRWMutex1000Read100Write(b *testing.B) {
	counter = 0
	var mu sync.RWMutex
	initGoroutines(1000, 100, mu.RLocker(), &mu)
}

func BenchmarkMutex1000Read1000Write(b *testing.B) {
	counter = 0
	var mu sync.Mutex
	initGoroutines(1000, 1000, &mu, &mu)
}

func BenchmarkRWMutex1000Read1000Write(b *testing.B) {
	counter = 0
	var mu sync.RWMutex
	initGoroutines(1000, 1000, mu.RLocker(), &mu)
}

func BenchmarkMutex1000Read10000Write(b *testing.B) {
	counter = 0
	var mu sync.Mutex
	initGoroutines(1000, 10000, &mu, &mu)
}

func BenchmarkRWMutex1000Read10000Write(b *testing.B) {
	counter = 0
	var mu sync.RWMutex
	initGoroutines(1000, 10000, mu.RLocker(), &mu)
}
