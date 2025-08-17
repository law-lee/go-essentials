package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getInts(chInts chan int) {
	chInts <- rand.Intn(1000)
}

func foo(ch chan int) {
	ch <- 1
	ch <- 2
	close(ch)
}

// You can close a channel with close(chan).

// Closing channel twice panics.

// Sending to closed channel panics.

// A receive from closed channels:

// returns buffered values
// return zero value immediately if there are no more buffered values
// Closing a channel also ends range loop over a channel.
func Run() {
	chInts := make(chan int)
	for i := 0; i < 2; i++ {
		go getInts(chInts)
	}

	n := <-chInts
	fmt.Printf("n: %d\n", n)

	// Using select allows to:
	// wait on multiple channels
	// do a non-blocking wait
	// implement timeouts by waiting on a timer channel
	select {
	case n := <-chInts:
		fmt.Printf("n: %d\n", n)
	}

	ch := make(chan int)
	go foo(ch)
	for n := range ch {
		fmt.Printf("value from foo chan is: %d\n", n)
	}

	// 	we use select to wait on 2 channels: chResult and a timeout channel
	// select finishes when receive on one of the 2 channels completes
	// we either get the value on chResult before timeout expires or we receive the value from timeout channel
	chResult := make(chan int, 1)

	go func() {
		time.Sleep(1 * time.Second)
		chResult <- 5
		fmt.Println("Worker finished")
	}()

	select {
	case res := <-chResult:
		fmt.Printf("Got %d from worker\n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out before worker finished")
	}

	chs := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for s := range chs {
			fmt.Printf("received from channel: %s\n", s)
		}
		fmt.Println("range loop finished because ch was closed")
		wg.Done()
	}()

	chs <- "string"
	close(chs)
	wg.Wait()

	unbuffered()
	buffered()

	chn, chQuit :=make(chan int), make(chan struct{})
	go worker(chn, chQuit)
	chn <- 3
	chQuit <- struct{}{}

}

func worker(ch chan int, chQuit chan struct{}) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value %d\n", v)
		case <-chQuit:
			fmt.Printf("Signalled on quit channel. Finishing\n")
			chQuit <- struct{}{}
			return
		}
	}
}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1 * time.Millisecond)
		}
		ch <- i
	}
}

func consumer(ch chan int) {
	total := 0
	for i := 0; i < 5; i++ {
		if i%2 == 1 {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1 * time.Millisecond)
		}
		total += <-ch
	}
}

func unbuffered() {
	timeStart := time.Now()
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
	fmt.Printf("Unbuffered version took %s\n", time.Since(timeStart))
}

func buffered() {
	timeStart := time.Now()
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
	fmt.Printf("Buffered version took %s\n", time.Since(timeStart))
}
