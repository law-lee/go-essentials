package main

import (
	"fmt"
	"time"
)

func TimerDemo() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("after 2s timeout!")

	res := timer.Stop()
	fmt.Println("stop timer after time out: ", res)

	timer.Reset(3 * time.Second)
	res = timer.Stop()
	fmt.Println("stop timer before time out: ", res)

	duration := time.Duration(1 * time.Second)
	f := func() {
		fmt.Println("f has been called after 1s by time.AfterFunc")
	}

	timer2 := time.AfterFunc(duration, f)
	defer timer2.Stop()
	time.Sleep(2 * time.Second)
}

func Watch() chan struct{} {
	ticker := time.NewTicker(1 * time.Second)
	ch := make(chan struct{})

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("watch!!!")
			case <-ch:
				fmt.Println("Ticker stop!!!")
				return
			}
		}
	}(ticker)
	return ch
}

func TickerDemo() {
	ch := Watch()
	time.Sleep(5 * time.Second)
	ch <- struct{}{}
	close(ch)
}

func main() {
	TimerDemo()
	TickerDemo()
}
