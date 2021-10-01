package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	m:=sync.Map{}
	wg:=sync.WaitGroup{}
	for i:=0;i<20;i++ {
		wg.Add(1)
		go func(n int) {
       		key:=strconv.Itoa(n)
       		m.Store(key,n)
       		value,_:=m.Load(key)
       		fmt.Printf("k=:%v v=:%v\n",key,value)
       		wg.Done()
        }(i)
	}
	wg.Wait()
}