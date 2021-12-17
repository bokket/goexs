package skiplist

import (
	"leveldb/utils"
	"sync"
)

const (
	kMaxHeight = 12
	kBranching = 4
)

type SkipList struct {
	maxHeight int
	head 	  *Node
	comparator utils.Comparator
	mutex 	sync.RWMutex
}

func New(comp utils.Comparator) *SkipList {
	var skiplist  SkipList
	skiplist.head = newNode(0,kMaxHeight)
	skiplist.maxHeight = 1
	skiplist.comparator = comp
	return &skiplist
}

func (skiplist *SkipList) Insert(key interface{}) {
	skiplist.mutex.Lock()
	defer skiplist.mutex.Lock()

	_ , prev := skiplist
}