package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Most of the time it’s not safe to access the same
// variable from multiple goroutines. Channels and sync.WaitGroup are exceptions.
var wg sync.WaitGroup

func sqrtWorker(chIn chan int, chOut chan int) {
	fmt.Println("sqrtWorker started")
	// Worker goroutines sqrtWorker pick up values from the channel using range.
	for i := range chIn {
		sqrt := i * i
		chOut <- sqrt
	}

	fmt.Println("sqrtWorker finished")
	// Just before terminating, the worker decrements wg counter.
	wg.Done()
}

func Run() {
	chIn := make(chan int)
	chOut := make(chan int)

	for i := 0; i < 2; i++ {
		// Before we launch the worker, we increment wg counter.
		wg.Add(1)
		go sqrtWorker(chIn, chOut)
	}

	go func() {
		//We don’t know which worker will pick any given value.
		chIn <- 2
		chIn <- 4
		close(chIn)
	}()

	go func() {
		wg.Wait()
		close(chOut)
	}()

	// There’s one more complication. Unless we close chOut, the for sqrt := range chOut loop will wait forever.
	for sqrt := range chOut {
		fmt.Printf("Got sqrt: %d\n", sqrt)
	}
}

var (
	semaphoreSize = 4

	mu                 sync.Mutex
	totalTasks         int
	curConcurrentTasks int
	maxConcurrentTasks int
)

func timeConsumingTask() {
	mu.Lock()
	totalTasks++
	curConcurrentTasks++
	if curConcurrentTasks > maxConcurrentTasks {
		maxConcurrentTasks = curConcurrentTasks
	}
	mu.Unlock()

	// in real system this would be a CPU intensive operation
	time.Sleep(10 * time.Millisecond)

	mu.Lock()
	curConcurrentTasks--
	mu.Unlock()
}

func Run2() {
	sem := make(chan struct{}, semaphoreSize)
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		// acquire semaphore
		sem <- struct{}{}
		wg.Add(1)

		go func() {
			timeConsumingTask()
			// release semaphore
			<-sem
			wg.Done()
		}()
	}

	// wait for all task to finish
	wg.Wait()

	fmt.Printf("total tasks         : %d\n", totalTasks)
	fmt.Printf("max concurrent tasks: %d\n", maxConcurrentTasks)
}
