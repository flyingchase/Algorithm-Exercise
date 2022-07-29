package main

import (
	"fmt"
)

func main() {
	num := 0
	defer func() {
		fmt.Println("num = ", num)
	}()
	// time.Sleep(1 * time.Second)
	num = 1
}
