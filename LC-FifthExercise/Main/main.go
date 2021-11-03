package main

import (
	"fmt"
	"time"
)

var primeCount int = 40

func PutInt(intChan chan<- int) {
	for i := 1; i < primeCount; i++ {
		intChan <- i
	}
	fmt.Println("put done")
	close(intChan)
}

func findPrime(intChan <-chan int, primeChan chan<- int, exitChan chan<- bool) {
	for {

		n, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(isPrime(n), n)
		if t := isPrime(n); t {
			primeChan <- n
		}
	}
	exitChan <- true
	fmt.Println("find Prime")
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var maxChan int = 4

func main() {
	intChan := make(chan int, primeCount)
	primeChan := make(chan int, primeCount)
	exitChan := make(chan bool, maxChan)
	go PutInt(intChan)

	go func() {

		for i := 0; i < maxChan; i++ {
			<-exitChan
		}
		fmt.Println("exit")
	}()

	for {
		time.Sleep(time.Second * 1)
		n, ok := <-primeChan
		if !ok {
			fmt.Println("end")
		}
		fmt.Println("prime = ", n, " ok = ", ok)
	}
}
