/**
 * @Author: lihaoran
 * @Date: 2022/7/4 22:03
 */

package main

type LRUCache struct {
	cap        int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// DLinkedNode 双向链表节点
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func (c *LRUCache) Put(k, v int) {
	// 存在 更新、移动到尾部
	if _, ok := c.cache[k]; ok {
		c.cache[k].value = v
		// 将节点移动到链表尾部
		c.moveNodeToTail(c.cache[k])
	} else {
		// 不存在，判断长度
		if len(c.cache) >= c.cap {
			delete(c.cache, c.head.next.key)
			// 删除头部节点
			c.deleteFromHead()
		}
		// 尾部插入
		node := &DLinkedNode{
			key:   k,
			value: v,
		}
		c.addToTail(node)
		c.cache[k] = node
	}
}

func (c *LRUCache) Get(k int) int {
	node, ok := c.cache[k]
	if !ok {
		return -1
	}
	// 将节点移动到链表尾部
	c.moveNodeToTail(node)
	return node.value
}

// Constructor 初始化设置头、尾节点
func Constructor(cap int) LRUCache {
	l := LRUCache{
		cap:   cap,
		cache: make(map[int]*DLinkedNode),
		head:  &DLinkedNode{},
		tail:  &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (c *LRUCache) moveNodeToTail(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	c.addToTail(node)
}

func (c *LRUCache) addToTail(node *DLinkedNode) {
	node.prev = c.tail.prev
	node.next = c.tail
	c.tail.prev.next = node
	c.tail.prev = node
}

func (c *LRUCache) deleteFromHead() {
	c.head.next = c.head.next.next
	c.head.next.prev = c.head
}
