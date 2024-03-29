package main

import (
	"context"
	"fmt"
	"time"
)

type result struct {
	data string
	err  error
}

func rpc() (string, error) {
	time.Sleep(100 * time.Millisecond)
	return "rpc done", nil
}
func handle(ctx context.Context, ms int) {
	ctx, cancle := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
	defer cancle()
	r := make(chan result)
	go func() {
		data, err := rpc()
		r <- result{data: data, err: err}
	}()
	select {
	case <-ctx.Done():
		fmt.Printf("timeout %d ms,context exit %+v\n", ms, ctx.Err())
	case res := <-r:
		fmt.Printf("res:%s, err %+v\n", res.data, res.err)
	}
}
func main() {
	// mock the requests goroutine to ask reqs
	for i := 1; i < 5; i++ {

		time.Sleep(1 * time.Second)
		go handle(context.Background(), i*50)
	}
	time.Sleep(time.Second)
}
