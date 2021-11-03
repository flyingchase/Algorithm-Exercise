package firstround

import (
	"container/list"
)

type (
	Pair struct {
		Key   int
		Value int
		Freq  int
	}
	LFUCache struct {
		cap    int
		freq   *list.List
		elemes map[int]*list.Element
	}
)

func Construcor(capacity int) LFUCache {
	return LFUCache{
		cap:    capacity,
		freq:   list.New(),
		elemes: make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	if elem, ok := this.elemes[key]; ok {
		this.touch(elem)
		return elem.Value.(*Pair).Value
	} else {
		return -1
	}
}

func (this *LFUCache) touch(elem *list.Element) {
	elem.Value.(*Pair).Freq += 1

	for e := elem.Prev(); e != nil; e = e.Prev() {
		if e.Value.(*Pair).Freq <= elem.Value.(*Pair).Freq {
			this.freq.MoveBefore(elem, e)
		}
	}
}

func (this *LFUCache) Put(key int, value int) {

	pair := Pair{
		Key:   key,
		Value: value,
	}
	if elem, ok := this.elemes[key]; ok {
		this.touch(elem)
		elem.Value.(*Pair).Value = value
	} else if len(this.elemes) >= this.cap {
		if this.cap == 0 {
			return
		}
		oldElem := this.freq.Back()
		delete(this.elemes, oldElem.Value.(*Pair).Key)
		this.freq.Remove(oldElem)
		elem := this.freq.PushBack(&pair)
		this.touch(elem)
		this.elemes[key] = elem
	} else {
		elem := this.freq.PushBack(&pair)
		this.touch(elem)
		this.elemes[key] = elem
	}

}
