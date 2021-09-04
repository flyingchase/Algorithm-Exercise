package HotReview
// list底层实现是双向链表
import "container/list"

/*
LRU 即为最近最少使用
更新和插入页面均在链表首 删除在链表尾

*/
type LRUCache struct {
    Cap int
    Keys map[int]*list.Element
    List *list.List

}

type pair struct {
    K,V int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        Cap: capacity,
        Keys: make(map[int] * list.Element),
        List: list.New(),
    }
}

type Element struct {
    next,prev *Element
    // list  存储的是头结点
    list *list.List

    value interface{}
}

func (c *LRUCache) Get(key int)  int{
    if el,ok:=c.Keys[key];ok {
        c.List.MoveToFront(el)
        return el.Value.(pair).V
    }
    return -1
}

func (c *LRUCache) Put(key int, value int)  {
    if el,ok:=c.Keys[key];ok {
        el.Value=pair{K:key,V: value}
        c.List.MoveToFront(el)
    }else {
        el:=c.List.PushFront(pair{K: key,V : value})
        c.Keys[key]=el
    }
    if c.List.Len() >c.Cap{
        el:=c.List.Back()
        c.List.Remove(el)
        delete(c.Keys,el.Value.(pair).K)
    }

}