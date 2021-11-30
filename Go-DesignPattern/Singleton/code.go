package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

var once sync.Once

func getInstance() *single {

	// if singleInstance == nil {
	// 	lock.Lock()
	// 	defer lock.Unlock()
	// 	if singleInstance == nil {
	// 		fmt.Println("creating single instance now")
	// 		singleInstance = &single{}
	// 	} else {
	// 		fmt.Println("single instance alreadly created-1")
	// 	}
	// } else {
	// 	fmt.Println("single instance alreadly created-2")
	// }
	// return singleInstance
	//

	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now")
				singleInstance = new(single)
			},
		)
	} else {
		fmt.Println("single instance alreadly created-2")
	}
	return singleInstance

}
func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
