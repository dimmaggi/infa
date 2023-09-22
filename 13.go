package main

import (
	"fmt"
	"sync"
)

type Node struct {
	index int
}

type NodeIndexer struct {
	counter int
	mutex   sync.Mutex
}

func (ni *NodeIndexer) GetNextIndex() int {
	ni.mutex.Lock()
	defer ni.mutex.Unlock()

	ni.counter++
	return ni.counter
}

func main() {
	indexer := NodeIndexer{}

	node1 := Node{index: indexer.GetNextIndex()}
	node2 := Node{index: indexer.GetNextIndex()}
	node3 := Node{index: indexer.GetNextIndex()}

	fmt.Println("Node 1 index:", node1.index)
	fmt.Println("Node 2 index:", node2.index)
	fmt.Println("Node 3 index:", node3.index)
}
