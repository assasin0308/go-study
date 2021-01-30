package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left,right *treeNode
}
// 工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setvalue(value int) {
	if node == nil {
		fmt.Println("set value nil to node ")
		return
	}
	node.value = value
}

func (node *treeNode) travel() {
	if node == nil {
		return
	}
	node.left.travel()
	node.print()
	node.right.travel()
}

func main() {
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	fmt.Println(root)

	nodes := []treeNode {
		{value: 3},
		{},
		{5,nil,&root},
	}
	fmt.Println(nodes)
	root.left.right = createNode(10)
	fmt.Println(createNode(10))
	root.print()
	root.right.left.setvalue(4)
	root.travel()
	root.right.left.print()

	var proot *treeNode
	proot.setvalue(200)
	proot = &root
	proot.setvalue(300)
	proot.print()

}
