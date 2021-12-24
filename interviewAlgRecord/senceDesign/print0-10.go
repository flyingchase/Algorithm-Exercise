package sencedesign

import (
	"fmt"
	"time"
)

/*
*并发编程
交替打印 1-10
* */
type Token struct{}

func AlternatelyPrint(total int) {
	channels := make([]chan Token, total)
	for i := 0; i < total; i++ {
		channels[i] = make(chan Token)
	}
	for i := 0; i < total; i++ {
		go func(index int, current chan Token, nextChan chan Token) {
			for {
				<-current
				fmt.Printf("Goroutine %d \n", index)
				time.Sleep(time.Second)
				nextChan <- Token{}
			}
		}(i+1, channels[i], channels[(i+1)%total])
	}
	channels[0] <- Token{}
	select {}
}
