package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	////stop:=make(chan bool)
	////go func() {
	////	for {
	////		select {
	////		case <-stop:
	////			fmt.Println("stop and quit")
	////			return
	////		default:
	////			fmt.Println("goroutine is under control")
	////			time.Sleep(2*time.Second)
	////		}
	////	}
	////}()
	////time.Sleep(5*time.Second)
	////fmt.Println(" now it is time to tell stop channel")
	////stop <- true
	////time.Sleep(3*time.Second)
	//
	//ctx,cancel:=context.WithCancel(context.Background())
	//go func(ctx context.Context) {
	//	for  {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("quit ")
	//			return
	//		default:
	//			fmt.Println("goroutine is under control")
	//			time.Sleep(2*time.Second)
	//		}
	//	}
	//}(ctx)
	//time.Sleep(10*time.Second)
	//fmt.Println("it is time to tell ctx stop control")
	//cancel()
	//time.Sleep(5*time.Second)

	gen := func(ctx context.Context) <-chan int {

		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	fmt.Println("________________________________")
	type favKey string
	f := func(ctx context.Context, k favKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value ", v)
			return
		}
		fmt.Println("not found key ", k)
	}

	k := favKey("language")
	ctx = context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, favKey("color"))

	fmt.Println("________________________________")

	ctx, cancle = context.WithTimeout(context.Background(), shortDuration)
	defer cancle()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

const shortDuration = 1 * time.Millisecond
