package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// fmt.Println(scope()()) // outpus:2
	// fmt.Println()
	// inner, val := outer()
	// fmt.Println(inner()) // outpus:200
	// fmt.Println(val)     // outpus:101
	// f("direct")
	// go f("goroutine")
	//
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")
	// time.Sleep(time.Second)
	// fmt.Println("Done")

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go w(i, &wg)
	}
	wg.Wait()
}

func w(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("%d done\n", id)
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func scope() func() int {
	outer_var := 2
	foo := func() int {
		return outer_var
	}
	return foo
}

func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99
		return outer_var
	}
	inner()
	return inner, outer_var
}
