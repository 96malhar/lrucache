package lrucache

type LRUCache struct {
	ll       *linkedList
	dict     map[int]*listNode
	capacity int
	size     int
}

type linkedList struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	key   int
	value int
	prev  *listNode
	next  *listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		ll:       &linkedList{},
		dict:     make(map[int]*listNode),
		capacity: capacity,
		size:     0,
	}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.dict[key]; !ok {
		return -1
	}
	node := c.dict[key]
	c.ll.delete(node)
	c.ll.append(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.dict[key]; ok {
		node := c.dict[key]
		node.value = value
		c.ll.delete(node)
		c.ll.append(node)
		return
	}

	if c.size == c.capacity {
		c.evictLRU()
		c.size--
	}
	node := createNode(key, value)
	c.dict[key] = node
	c.ll.append(node)
	c.size++
}

func (c *LRUCache) evictLRU() {
	lruNode := c.ll.head
	if lruNode == nil {
		return
	}

	c.ll.delete(lruNode)
	delete(c.dict, lruNode.key)
}

func (ll *linkedList) append(node *listNode) {
	if node == nil {
		return
	}
	if ll.head == nil {
		ll.head = node
	} else {
		ll.tail.next = node
		node.prev = ll.tail
	}
	ll.tail = node
}

func (ll *linkedList) delete(node *listNode) {
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node == ll.head {
		ll.head = node.next
	}
	if node == ll.tail {
		ll.tail = node.prev
	}
	node.next = nil
	node.prev = nil
}

func createNode(key, value int) *listNode {
	return &listNode{
		key:   key,
		value: value,
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
