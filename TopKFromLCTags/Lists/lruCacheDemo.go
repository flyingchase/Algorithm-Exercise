package topkLists

type Node struct {
	Next, Prev *Node
	list       *lru
	key        string
	value      interface{}
}
type lru struct {
	len  int
	root *Node
	cap  int
}

func (l *lru) get() *Node {
}
