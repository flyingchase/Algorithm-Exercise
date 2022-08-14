package HotReview

import (
	"sync"
)

type LRUCache struct {
	head, tail *Node
	cap, len   int
	keys       map[int]*Node
	mu         *sync.RWMutex
}
type Node struct {
	prev, next *Node
	key        int
	val        interface{}
}

func Constructor(capability int) LRUCache {
	return LRUCache{
		keys: make(map[int]*Node),
		mu:   new(sync.RWMutex),
		cap:  capability,
	}
}

func (this *LRUCache) Put(key int) interface{} {
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.cap <= 0 {
		return nil
	}
	// find in the keys map
	if node, ok := this.keys[key]; ok {
		this.remove(node)
		this.add(node)
		return node.val
	}
	return nil
}

func (this *LRUCache) Push(key int, value interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()

	// already exists in keysMap
	if node, ok := this.keys[key]; ok {
		node.val = value
		this.remove(node)
		this.add(node)
		return
	}

	// constructor new node to add
	node := &Node{
		key: key,
		val: value,
	}
	this.add(node)
	return
}

func (this *LRUCache) add(node *Node) {

	node.next = this.head
	node.prev = nil
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}

	this.len++

	if this.cap <= this.len {
		delete(this.keys, this.tail.key)
		this.remove(this.tail)
		this.len--
	}

}

func (this *LRUCache) remove(node *Node) {
	if this.head == node {
		this.head = node.next
		if this.head != nil {
			this.head.prev = nil
		}
		this.len--
		node.prev = nil
		node.next = nil
		return
	}

	if this.tail == node {
		this.tail = node.prev
		if this.tail != nil {
			this.tail.next = nil
		}
		this.len--
		node.prev = nil
		node.next = nil
		return
	}

	this.len--

	node.prev.next = node.next
	node.next.prev = node.prev
}

func findCommonSubString(word1, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	var res int

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				if i >= 1 && j >= 1 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = 0
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

func findSearchInRoateNums(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[l] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] == nums[l] {
				l++
			}
			if nums[mid] == nums[r] {
				r--
			}
		}

	}
	return -1
}

func lis(str string) int {

	return -1
}
