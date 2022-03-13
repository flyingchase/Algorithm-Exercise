package implementions

// 实现双向链表
// 法一：使用 head 和 tail 指针指向首尾结点
// 法二：使用 head 和 tail 作为头尾虚拟节点，不存储数据；统一 crud 操作
type DoubleListNode struct {
	Prev, Next *DoubleListNode
	Val        int
}

var head, tail = &DoubleListNode{}, &DoubleListNode{}

func NewDoubleListNode() *DoubleListNode {
	head.Next = tail
	tail.Prev = head

	return &DoubleListNode{
		Prev: head,
		Next: tail,
	}
}
func (d *DoubleListNode) Get(index int) *DoubleListNode {
	for i := 0; i < index; i++ {
	}
	return &DoubleListNode{}
}
func (d *DoubleListNode) AddAtHead(val int) bool {
	node := &DoubleListNode{Val: val}
	next := d.Next

	head.Next = node
	node.Prev = head
	node.Next = next
	next.Prev = node
	return true
}
func (d *DoubleListNode) AddAtTail(val int) bool {
	node := &DoubleListNode{Val: val}

	prev := tail.Prev
	node.Next = tail
	node.Prev = prev
	tail.Prev = node
	prev.Next = node
	return true
}
func (d *DoubleListNode) AddAtIndex(index, val int) bool {
	return true

}
func (d *DoubleListNode) DeleteAtIndex(index int) {

}
