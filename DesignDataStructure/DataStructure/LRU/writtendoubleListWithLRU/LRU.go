package writtendoublelistwithlru

import (
	"fmt"
)

type (
	Node struct {
		prev, next *Node
		list       *LRU
		key        string
		value      interface{}
	}
	LRU struct {
		root *Node
		cap  int
		len  int
	}
)

// 双向循环链表，表头记为 root 不存元素作为 dummyHead 使用
// 查询元素，调至队首
// 弹出元素，判断是否存在 key ，判断是否超容，再插入链首
func NewLRU(cap int) *LRU {
	l := &LRU{
		cap:  cap,
		root: &Node{},
	}
	l.root.prev = l.root
	l.root.next = l.root
	l.root.list = l
	return l
}
func (l *LRU) get(key string) *Node {
	defer l.debug()

	for node := l.root.next; node != l.root; node = node.next {
		if node.key == key {
			// 断开 node 前后结点并连接上前后结点
			node.prev.next = node.next
			node.next.prev = node.prev
			// 找到后移至队首
			node.next = l.root.next
			l.root.next.prev = node
			l.root.next = node
			node.prev = l.root
			return node
		}
	}
	return nil
}
func (l *LRU) debug() {
	fmt.Println("lru len : ", l.len)
	fmt.Println("lru cap : ", l.cap)
	for node := l.root.next; node != l.root; node = node.next {
		fmt.Printf("%s : %v ", node.key, node.value)
	}
	fmt.Println()
}

func (l *LRU) Get(key string) interface{} {
	defer l.debug()
	n := l.get(key)
	if n == nil {
		return nil
	}
	return n.value
}

func (l *LRU) Put(key string, value interface{}) {
	defer l.debug()
	n := l.get(key)
	// key 已存在则更新 value
	if n != nil {
		n.value = value
		return
	}
	// key不存在且 cap 满
	// delete the last node and l.len--
	if l.len == l.cap {
		last := l.root.prev
		last.prev.next = l.root
		l.root.prev = last.prev

		last.prev = nil
		last.next = nil
		last.list = nil
		l.len--
	}
	// 构造新节点插入队首，更新 l.len 和新节点的 list 所属
	node := &Node{
		key:   key,
		value: value,
	}
	head := l.root.next
	head.prev = node
	node.next = head
	node.prev = l.root
	l.root.next = node
	l.len++
	node.list = l
}
