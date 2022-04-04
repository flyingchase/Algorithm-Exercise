package innerdoublelistswithlfu

import (
	"container/list"
)

type LFUCache struct {
	nodes map[int]*list.Element
	// 相同频次的 nodes 的先后顺序排列，多个相同频次则多个双向链表
	lists    map[int]*list.List
	capacity int
	// 最少访问
	min int
}
type node struct {
	key   int
	value int
	// 访问频次
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.nodes[key]
	// 不存在 key 对应的 value 则-1
	if !ok {
		return -1
	}
	curNode := value.Value.(*node)
	// 删除当前 node 的结点并将 frequency+1
	this.lists[curNode.frequency].Remove(value)
	curNode.frequency++
	// 存在增后的 frequency,则移至链首,不存在则创建新链表并移至队首
	if _, ok := this.lists[curNode.frequency]; ok {
		this.lists[curNode.frequency] = list.New()
	}
	newList := this.lists[curNode.frequency]
	newNode := newList.PushFront(curNode)
	this.nodes[key] = newNode
	// 更新 min 值，原frequency对应的 lists 为空则++
	if curNode.frequency-1 == this.min && this.lists[curNode.frequency-1].Len() == 0 {
		this.min++
	}
	return curNode.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// 结点存在则更新访问次数，移动结点
	if curValue, ok := this.nodes[key]; ok {
		curNode := curValue.Value.(*node)
		curNode.value = value
		this.Get(key)
		return
	}
	// 不存在结点且缓存满，删除再添加
	if this.capacity == len(this.nodes) {
		curList := this.lists[this.min]
		backNode := curList.Back()
		delete(this.nodes, backNode.Value.(*node).key)
		curList.Remove(backNode)
	}
	this.min = 1
	curNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := this.lists[1]; ok {
		this.lists[1] = list.New()
	}
	newList := this.lists[1]
	newNode := newList.PushFront(curNode)
	this.nodes[key] = newNode
}
