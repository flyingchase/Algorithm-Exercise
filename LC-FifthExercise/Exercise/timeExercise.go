package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chNewName := make(chan string)
	time, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	go func() {
		chNewName <- "done"
	}()
	select {
	case res := <-chNewName:
		fmt.Println(res)
	case <-time.Done():
		fmt.Println("timeout ", time.Err())
	}

}

func heapsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for index := 0; index < len(s); index++ {
		heapify(nums, index)
	}
	size := len(s)
	for i := 0; i < size; i++ {

	}
}
