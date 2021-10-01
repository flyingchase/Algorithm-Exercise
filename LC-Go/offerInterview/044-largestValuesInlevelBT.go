package offerInterview

import (
	"LC-Go/DataStructure"
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type TreeNode = DataStructure.TreeNode

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		maxElement := math.MinInt32
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxElement {
				maxElement = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if cur.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		res = append(res, maxElement)
	}
	return res
}

type Pair struct {
	value int
	count int
}
type PriorityQueue []Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count< pq[j].count
}
func (pq PriorityQueue)Swap(i,j int)  {
	pq[i],pq[j]=pq[j],pq[i]
}
func (pq *PriorityQueue) Push(x interface{})  {
	*pq = append(*pq,x.(Pair))
}
func (pq *PriorityQueue)Pop() interface{}{
	old:=*pq
	top:=old[len(old)-1]
	*pq=old[: len(old)-1]
	return top
}
func topKthfrequenctInNums(nums []int,k int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})
	topK:=nums[k-1]
	fmt.Println(topK)
	m:=map[int]int{}
	for _,v:=range nums{
		m[v]+=1
	}

	pq:=&PriorityQueue{}
	heap.Init(pq)
	for val,cout:=range m{
		heap.Push(pq,Pair{val,cout})
	}
	res:=make([]int,0)
	for i:=0;i<k;i++{
		res=append(res,heap.Pop(pq).(Pair).value)
	}
	return res
}
