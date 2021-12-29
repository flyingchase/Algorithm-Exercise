package doublelistwithlru

import "container/list"

type LRUCache struct {
	Cap int
	// Element 为双向结点,内含next prev 指针
	// list 链表指针，Valure interface{}接口
	Keys map[int]*list.Element
	List *list.List
}

// list.Element存储的value存储
type pair struct {
	K, V int
}

// 利用 cap 创建 LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

// 通过 key 查找 value,不存在则返回-1
// 查找完更新 element 到链首
func (c *LRUCache) Get(key int) int {
	if el, ok := c.Keys[key]; ok {
		c.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

// k-v 来 Put 操作
// 先查询 map 是否存在 key，存在则更新 value 并移至链首
// 不存在则新建结点并插入双向链表和 map 中
// 维护 cap，超容则淘汰链尾结点并删除 map 对应的 key
func (c *LRUCache) Put(key int, value int) {
	if el, ok := c.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		c.List.MoveToFront(el)
	} else {
		el := c.List.PushFront(pair{K: key, V: value})
		c.Keys[key] = el
	}
	// 维护 cap
	if c.List.Len() > c.Cap {
		el := c.List.Back()
		c.List.Remove(el)
		delete(c.Keys, el.Value.(pair).K)
	}
}
