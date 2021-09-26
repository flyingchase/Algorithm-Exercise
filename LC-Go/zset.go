package main

const (
	maxLevel    = 32
	probability = 0.25
)

type (
	sklLevel struct {
		forward *sklNode
		span    uint64
	}
	sklNode struct {
		member   string
		score    float64
		backward *sklNode
		level    []*sklLevel
	}
	skipList struct {
		head   *sklNode
		tail   *sklNode
		length int64
		level  int16
	}
	SortedSetNode struct {
		dict map[string]*sklNode
		skl  *skipList
	}
	SortedSet struct {
		record map[string]*SortedSetNode
	}
)

func New() *SortedSet {
	return &SortedSet{
		make(map[string]*SortedSetNode),
	}


}
func (z *SortedSet)exist(key string)bool  {
	_,exist:=z.record[key]
	return exist

}
func (z *SortedSet) ZAdd(key string, score float64, member string) {
	if !z.exist(key) {
		node:=&SortedSetNode {
		    dict: make(map[string] *sklNode),
		    skl: newSkipList(),
		}

		z.record[key]=node
	}
	item:=z.record[key]
	v,exist:=item.dict[member]
}

func newSkipList() *skipList {
	return & skipList {
	    level: 1,
	    head: sklNewNode(maxLevel,0,""),
	}

}

func sklNewNode(level int16, score float64, member string) *sklNode {
	node := & sklNode {
		score: score,
		member: member,
		level: make([]*sklLevel,level),
	}
	for i:=range node.level{
		node.level[i]=new(sklLevel)
	}
	return node
}
