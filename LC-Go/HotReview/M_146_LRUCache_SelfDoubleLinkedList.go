package HotReview

type Node struct {
	key,val int
	prev,next *Node
}

type LRUCacheSelf struct {
	head,tail *Node
	keys map[int]*Node
	capacity int

}

func ConstructorLRU(capacity int) LRUCacheSelf {
	return LRUCacheSelf {
	    keys: make(map[int] * Node),
		capacity: capacity,
	}
}
func (this *LRUCacheSelf) Get(Key int) int {
	if node,ok:=this.keys[Key];ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCacheSelf) Remove(node *Node) {
	if node == this.head {
		this.head=node.next
		if node.next != nil {
			node.next.prev=nil
		}
		node.next=nil
		return
	}
	if node == this.tail {
		this.tail=node.prev
		node.prev.next=nil
		node.prev=nil
		return
	}

	node.prev.next=node.next
	node.next.prev=node.prev
}

func (this *LRUCacheSelf) Add(node *Node) {

	node.prev=nil
	node.next=this.head
	if this.head!=nil {
		this.head.prev=node
	}
	this.head=node
	if this.tail==nil {
		this.tail=node
		this.tail.next=nil
	}
}