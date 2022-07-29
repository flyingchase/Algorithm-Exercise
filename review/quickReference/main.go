package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("Done")
	}()
	fmt.Println("working")
	// time.Sleep(time.Second)
	var d = int64(0)
	defer func(d *int64) {
		fmt.Printf("& %v Unix Sec\n", *d)
	}(&d)
	fmt.Println("Done")
	d = time.Now().Unix()
}
