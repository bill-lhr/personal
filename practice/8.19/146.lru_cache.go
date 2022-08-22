package __19

type LRUCache struct {
	cache      map[int]*DLinkedNode
	cap        int
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, Val   int
	prev, next *DLinkedNode
}

func NewLRUCache(cap int) LRUCache {
	l := LRUCache{
		cache: make(map[int]*DLinkedNode, cap),
		cap:   cap,
		head:  &DLinkedNode{},
		tail:  &DLinkedNode{},
	}
	l.head.next, l.tail.prev = l.tail, l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	l.MoveToHead(node)
	return node.Val
}

func (l *LRUCache) Put(key int, val int) {
	node, ok := l.cache[key]
	if ok {
		node.Val = val
		l.MoveToHead(node)
		return
	}
	newNode := &DLinkedNode{
		key: key,
		Val: val,
	}
	if len(l.cache) == l.cap {
		l.DelFromTail()
	}
	l.AddToHead(newNode)
}

func (l *LRUCache) AddToHead(node *DLinkedNode) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next = node
	node.next.prev = node
	l.cache[node.key] = node
}

func (l *LRUCache) DelFromTail() {
	delete(l.cache, l.tail.prev.key)
	l.tail.prev = l.tail.prev.prev
	l.tail.prev.next = l.tail
}

func (l *LRUCache) MoveToHead(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
	l.AddToHead(node)
}
