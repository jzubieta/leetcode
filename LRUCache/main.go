package main

func (c *LRUCache) Get(key int) int

func (c *LRUCache) Put(key int, value int)

type Node struct {
    // Define tus campos
	key int
	value int
	prev *Node
	next *Node
}

type LRUCache struct {
   mostUsed *Node
   lessUsed *Node
   nodes map[int]*Node
   capacity int
}

func Constructor(capacity int) *LRUCache {
    // Implementar
	return &LRUCache{capacity: capacity,nodes: make(map[int]*Node,capacity)}
}

func (c *LRUCache) Get(key int) int {
    // Implementar
}

func (c *LRUCache) Put(key int, value int) {
    // Implementar
}