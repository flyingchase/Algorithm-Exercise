package implementions

import (
	"errors"
	"fmt"
	"sync"
)

type Queue interface {
	Size() int
	IsEmpty() bool
	IsFull() bool
	EnQueue(value interface{}) error
	DeQueue() (interface{}, error)
	GetQueueFront() (interface{}, error)
	GetQueueTail() (interface{}, error)
	PrintAll()
}

const DefaultRingQueueSize int = 5

type RingQueue struct {
	data  []interface{}
	front int
	tail  int
	lock  sync.Mutex
}

func NewRingQueue() *RingQueue {
	return &RingQueue{
		data: make([]interface{}, DefaultRingQueueSize, DefaultRingQueueSize),
	}
}
func (r *RingQueue) Size() int {
	r.lock.Lock()
	defer r.lock.Unlock()
	return (r.tail-r.front)%DefaultRingQueueSize + 1
}
func (r *RingQueue) IsFull() bool {
	return (r.tail+1)&DefaultRingQueueSize == r.front
}
func (r *RingQueue) IsEmpty() bool {
	return r.tail == r.front
}
func (r *RingQueue) EnQueue(value interface{}) error {
	if r.IsFull() {
		return errors.New("full Queue")
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	r.data[r.tail] = value
	r.tail = (r.tail + 1) & DefaultRingQueueSize
	return nil
}
func (r *RingQueue) DeQueue() (interface{}, error) {
	if r.IsEmpty() {
		return nil, errors.New("empty Queue")
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	value := r.data[r.front]
	r.front = (r.front + 1) & DefaultRingQueueSize
	return value, nil
}
func (r *RingQueue) GetQueueTail() (interface{}, error) {
	if r.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	r.lock.Lock()
	defer r.lock.Unlock()

	value := r.data[(r.tail-1+DefaultRingQueueSize)%DefaultRingQueueSize]
	return value, nil
}
func (r *RingQueue) GetQueueFront() (interface{}, error) {
	if r.IsEmpty() {
		return nil, errors.New("empty Queue")
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.data[r.front], nil
}

func (r *RingQueue) PrintAll() {
	if r.tail < r.front {
		for i := r.front; i < DefaultRingQueueSize; i++ {
			fmt.Print(r.data[i], "<-")
		}
		for i := 0; i < r.tail; i++ {
			if i == r.tail-1 {
				fmt.Print(r.data[i], "\n")
			} else {
				fmt.Print(r.data[i], "<-")
			}
		}
	} else {
		for i := r.front; i < r.tail; i++ {
			if i == r.tail-1 {
				fmt.Print(r.data[i], "\n")
			} else {
				fmt.Print(r.data[i], "<-")
			}
		}
	}
}
