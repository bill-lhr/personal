/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/4 11:08 下午
 */

package main

type Object interface {
}

// 链表结点结构
type Node struct {
	Data Object
	Next *Node
}

// 链表结构
// 链表本身是一个指向第一个结点的指针，这里定义成了一个struct
type LinkedList struct {
	Head *Node
}

// 判断是否为空的单链表，list为头指针，指向第一个结点
func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

// 单链表长度
func (list *LinkedList) Length() int {
	curNode, count := list.Head, 0
	for curNode != nil {
		count++
		curNode = curNode.Next
	}
	return count
}

// 获取第一个结点
func (list *LinkedList) GetFirstNode() *Node {
	return list.Head
}

// 从头部添加元素
func (list *LinkedList) Add(data Object) {
	node := &Node{
		Data: data,
	}

	node.Next = list.Head
	list.Head = node
}

// 从尾部添加元素
func (list *LinkedList) Append(data Object) {
	node := &Node{
		Data: data,
	}
	if list.IsEmpty() {
		list.Head = node
	} else {
		tail := list.Head
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = node
	}
}

// 指定位置添加元素，index表示插入位置，第一个结点规定为0
func (list *LinkedList) Insert(index int, data Object) {
	if index <= 0 {
		list.Add(data)
	} else if index > list.Length() {
		list.Append(data)
	} else {
		node := &Node{
			Data: data,
		}
		pre, curIndex := list.Head, 1
		for curIndex < index {
			pre = pre.Next
			curIndex += 1
		}
		node.Next = pre.Next
		pre.Next = node
	}
}

// 删除指定元素
func (list *LinkedList) Remove(data Object) {
	pre := list.Head

	if pre == data {
		list.Head = pre.Next
	} else {
		// 删除第一个指定元素
		//for pre.Next != nil && pre.Next.Data != data {
		//	pre = pre.Next
		//}
		//if pre.Next != nil {
		//	pre.Next = pre.Next.Next
		//}
		// 删除所有的指定元素
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// 删除指定位置的元素
func (list *LinkedList) RemoveIndex(index int) {
	pre := list.Head

	if index <= 0 {
		list.Head = pre.Next
	} else if index > list.Length() {
		// 错误
		return
	} else {
		preIndex := 1
		for preIndex < index {
			preIndex += 1
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// 是否包含某个元素
func (list *LinkedList) Contains(data Object) bool {
	cur := list.Head

	for cur.Next != nil && cur.Data != data {
		cur = cur.Next
	}
	return cur.Data == data
}
