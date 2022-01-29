package innerlistswithlock

import (
	"container/list"
	"sync"
)

type LruCache struct {
	size     int
	values   *list.List
	cacheMap map[interface{}]*list.Element
	lock     sync.Mutex
}

func NewLruList(size int) *LruCache {
	value := list.New()
	return &LruCache{
		size:     size,
		values:   value,
		cacheMap: make(map[interface{}]*list.Element, size),
	}
}
func (l *LruCache) Put(k, v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.values.Len() == l.size {
		back := l.values.Back()
		l.values.Remove(back)
		delete(l.cacheMap, k)
	}
	front := l.values.PushFront(v)
	l.cacheMap[k] = front
}
func (l *LruCache) Get(k interface{}) (interface{}, bool) {
	v, ok := l.cacheMap[k]
	if ok {
		l.values.MoveToFront(v)
		return v, true
	}
	return nil, false
}
