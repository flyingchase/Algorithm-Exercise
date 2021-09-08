package main

import (
	"fmt"
	"time"
)

func main() {
	notify := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				select {
				case <-notify:
					fmt.Println("done...", i)
				case <-time.After(1 * time.Second):
					fmt.Println("wait notify", i)
				}
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	close(notify)
	time.Sleep(3 * time.Second)
}
