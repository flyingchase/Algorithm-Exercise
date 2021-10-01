package main

import (
	"fmt"
	"sync"
)

func hello()  {

	fmt.Println("now goroutine is hello func")
}
func main() {
	for i:=0;i<10;i++ {
		wg.Add(1)
		go hellotwo(i)
	}
	wg.Wait()
}

var wg sync.WaitGroup

func hellotwo(i int)  {
	defer wg.Done()
	fmt.Println("hello goroutine ",i)
}
