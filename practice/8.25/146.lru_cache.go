package __25

type LRUCache struct {
	cache      map[int]*DLinkedList
	cap        int
	head, tail *DLinkedList
}

type DLinkedList struct {
	key, val   int
	prev, next *DLinkedList
}

func NewLRUCache(cap int) *LRUCache {
	cache := &LRUCache{
		cache: make(map[int]*DLinkedList),
		cap:   cap,
		head:  &DLinkedList{},
		tail:  &DLinkedList{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (l *LRUCache) Get(k int) int {
	node, ok := l.cache[k]
	if !ok {
		return -1
	}
	l.MoveToHead(node)
	return node.val
}

func (l *LRUCache) Put(k, v int) {
	node, ok := l.cache[k]
	if ok {
		node.val = v
		l.MoveToHead(node)
		return
	}
	if len(l.cache) == l.cap {
		l.DeleteFromTail()
	}
	l.AddToHead(&DLinkedList{
		key: k,
		val: v,
	})
}

func (l *LRUCache) AddToHead(node *DLinkedList) {
	node.next = l.head.next
	node.next.prev = node
	l.head.next = node
	node.prev = l.head
	l.cache[node.key] = node
}

func (l *LRUCache) MoveToHead(node *DLinkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.AddToHead(node)
}

func (l *LRUCache) DeleteFromTail() {
	node := l.tail.prev
	node.prev.next = l.tail
	l.tail.prev = node.prev
	delete(l.cache, node.key)
}
