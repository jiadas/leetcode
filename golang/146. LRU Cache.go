package golang

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func NewDLinkNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: val,
	}
}

func LRUCacheConstructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode, capacity+1),
		head:     NewDLinkNode(0, 0),
		tail:     NewDLinkNode(0, 0),
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		node.value = value
		return
	}
	node := NewDLinkNode(key, value)
	c.cache[key] = node
	c.addToHead(node)
	c.size++
	if c.size > c.capacity {
		removed := c.removeTail()
		delete(c.cache, removed.key)
		c.size--
	}
}

func (c *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) moveToHead(node *DLinkedNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *DLinkedNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
