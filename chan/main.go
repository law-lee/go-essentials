package main

import (
	"fmt"
	"time"
)

func add(ch chan bool, num *int) {
	ch <- true
	*num = *num + 1
	<-ch
}

func main() {
	// 创建一个size为1的channel
	ch := make(chan bool, 1)

	var num int
	for i := 0; i < 100; i++ {
		go add(ch, &num)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("num 的值：", num)
}
