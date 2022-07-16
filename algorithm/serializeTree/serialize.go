/**
 * @Author: lihaoran04@baidu.com
 * @Date: 2020/8/31 11:46 下午
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strList := strings.Split(data, ",")
	return dfs2(&strList)
}

func dfs2(data *[]string) *TreeNode {
	strval := (*data)[0]
	*data = (*data)[1:]
	if strval == "#" {
		return nil
	}

	val, _ := strconv.Atoi(strval)
	node := &TreeNode{Val: val}

	node.Left = dfs2(data)
	node.Right = dfs2(data)
	return node
}

func main() {
	root := &TreeNode{}

	obj := Constructor()
	data := obj.serialize(root)
	ans := obj.deserialize(data)
	fmt.Println(ans)
}
