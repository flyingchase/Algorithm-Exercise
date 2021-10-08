// // /condition*
//   * @Author: wlzhou
//  * @Date:   2021-09-06 11:50:59
//  * @Last Modified time: 2021-09-06 12:27:31
//  */
// // package main
package exercises

import (
	"fmt"
	"testing"
	"time"
)

func testAppendA(t *testing.T) {
	x := []int{1, 2, 3}
	appendA(x)
	fmt.Printf("main %v\n", x)
}

func appendA(x []int) {
	x[0] = 100
	fmt.Printf("appendA %v\n", x)
}

func main() {
	ch := make(chan struct{})

	go func() {
		for i := 1; i < 11; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Println("奇数", i)
			}
		}
	}()
	go func() {

		for i := 0; i < 11; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("偶数", i)
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
