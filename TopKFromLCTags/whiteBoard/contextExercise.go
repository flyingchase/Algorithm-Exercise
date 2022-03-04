package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("quit ", name)
			return
		default:
			fmt.Println("send request ", name)
			time.Sleep(time.Second)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx, "1")
	go reqTask(ctx, "2")
	go reqTask(ctx, "3")

	time.Sleep(3 * time.Second)
	fmt.Println("it is time to tell cancel stop control")
	cancel()
	time.Sleep(3 * time.Second)
}
