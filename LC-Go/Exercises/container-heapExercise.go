package main

import (
	"container/heap"
	"context"
	"fmt"
	"time"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := * h
	n := len(old)
	*h = old[:n-1]
	x := old[n-1]
	return x
}

type Item struct {
	value string
	// priority means the order of item in queue
	priority int
	// index means the order of item in heap
	index int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(* pq, item)
	}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	x.index = -1
	*pq = old[:n-1]
	return x
}
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	//l := list.New()
	//e4 := l.PushBack(4)
	//e1 := l.PushFront(1)
	//l.InsertAfter(2, e1)
	//l.InsertBefore(3, e4)
	//for e := l.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	//var wg sync.WaitGroup
	//
	//wg.Add(2)
	//
	//go func(){
	//	time.Sleep(2*time.Second)
	//	fmt.Println("first Done")
	//	wg.Done()
	//}()
	//
	//
	//go func() {
	//	time.Sleep(2*time.Second)
	//	fmt.Println("second Done")
	//	wg.Done()
	//}()
	//wg.Wait()
	//pq:=&PriorityQueue{}
	//heap.Init(pq)
	//m:=make(map[int]int)
	//nums:=[]int{12312312,213,2131}
	//for _,v:=range nums{
	//	m[v]+=1
	//}
	//for k,v:=range nums{
	//	heap.Push(pq,&Item{
	//		index: 0,
	//		priority: v,
	//		value: string(k),
	//	})
	//}
	//
	//
	//fmt.Println("all Done")
ctx,cancle:=context.WithCancel(context.Background())
go reqTask(ctx,"one")
go reqTask(ctx,"two")
go reqTask(ctx,"three")

time.Sleep(3*time.Second)
cancle()
time.Sleep(3*time.Second)
}
func reqTask(ctx context.Context,name string)  {
	for  {
		select {
		case <-ctx.Done():
			fmt.Println("stop",name)
			return
		default:
			fmt.Println("send request: ",name)
			time.Sleep(time.Second)
		}

	}

}

//h := &IntHeap{2, 5, 6, 0, 1, 8, 3, 9, 10, 7}
//heap.Init(h)
//heap.Push(h, -1)
//minElemnt := (*h)[0]
//fmt.Printf("minheap top is %d\n", minElemnt)
//
//
//items:=map[string]int{
//	"one":1,
//	"two":2,
//	"three":3,
//}
//
//pq:=make(PriorityQueue,len(items))
//i:=0
//for v,priority:=range items{
//	pq[i]=&Item{
//		value: v,
//		priority: priority,
//		index: i,
//	}
//	i++
//}
//heap.Init(&pq)
//
//item:=&Item{
//	value: "ONE",
//	priority: 1,
//}
//heap.Push(&pq,item)
//pq.update(item,item.value,2)
//
//for pq.Len()>0 {
//	item:=heap.Pop(&pq).(*Item)
//	fmt.Printf("%.2d:%s\n",item.priority,item.value)
//
//}
