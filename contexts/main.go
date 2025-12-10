package main

import (
	"context"
	"fmt"
	"time"
)

// context 的作用
// 1. 用于并发控制，控制携程的优雅退出，比如在goroutine 之间传递取消信号\超时信号
// 2. 用于传递一些请求级别的数据，比如用户 ID、请求 ID 等等

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(4*time.Second)) // 设置超时时间4当前时间4s之后

	go Watch(ctx, "goroutine1")
	go Watch(ctx, "goroutine2")

	time.Sleep(6 * time.Second)
	fmt.Println("end watching!!!")
	cancel()
	time.Sleep(1 * time.Second)
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s exit!\n", name)
			return
		default:
			fmt.Printf("%s watching...\n", name)
			time.Sleep(time.Second)
		}
	}
}

// func func1(ctx context.Context) {
//    fmt.Printf("name is: %s", ctx.Value("name").(string))
// }

// func main() {
//    ctx := context.WithValue(context.Background(), "name", "zhangsan")
//    go func1(ctx)
//    time.Sleep(time.Second)
// }
