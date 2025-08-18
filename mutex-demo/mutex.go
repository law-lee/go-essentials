package mutexdemo

import (
	"fmt"
	"sync"
	"time"
)

var cache map[int]int

// A sync.RWMutex has 2 types of lock function: lock for reading and lock for writing.
// It follows the following rules: * a writer lock takes exclusive lock * a reader
// lock will allow another readers but not writer
// var mu sync.Mutex
var mu sync.RWMutex

func expensiveOperation(n int) int {
	// in real code this operation would be very expensive
	return n * n
}

func getCached(n int) int {
	// 	In some languages mutexes are recursive i.e. the same thread can Lock the same mutex multiple times.
	// In Go sync.Mutex is non-recursive. Calling Lock twice in the same goroutine will deadlock.
	mu.RLock()
	v, isCached := cache[n]
	mu.RUnlock()
	if isCached {
		return v
	}

	v = expensiveOperation(n)

	mu.Lock()
	cache[n] = v
	mu.Unlock()
	return v
}

func accessCache() {
	now := time.Now()
	total := 0
	for i := 0; i < 5; i++ {
		n := getCached(i)
		total += n
	}
	fmt.Printf("total: %d, time elapsed: %d\n", total, time.Since(now).Milliseconds())
}

func Run() {
	cache = make(map[int]int)
	go accessCache()
	accessCache()
}
