package main

import (
	"fmt"
)

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}

func NewBinaryNode(val int) *BinaryNode {
	return &BinaryNode{val, new(BinaryNode), new(BinaryNode)}
}

func (n *BinaryTree) Add(val int) {
	if n.Root == nil {
		n.Root = NewBinaryNode(val)
	} else {
		n.Root.Add(val)
	}
}

func (n *BinaryNode) Add(val int) {
	if val < n.Value {
		if n.Left != nil {
			n.Left.Add(val)
		} else {
			n.Left = NewBinaryNode(val)
		}
	} else {
		if n.Right != nil {
			n.Right.Add(val)
		} else {
			n.Right = NewBinaryNode(val)
		}
	}
}

func (n *BinaryTree) Contains(target int) bool {
	node := n.Root
	for node != nil {
		if target > node.Value {
			node = node.Right
		} else if target < node.Value {
			node = node.Left
		} else {
			return true
		}
	}
	return false
}

func mains() {
	dict := make(map[string]string)
	dict["apple"] = "120円"
	dict["orange"] = "90円"
	for key, value := range dict {
		fmt.Println(key)
		fmt.Println(value)
	}
	ary := []string{}
	slice := []int{}
	ary = append(ary, "apple")
	for i := 0; i < 10; i++ {
		slice = append([]int{i}, slice...)
	}
	fmt.Println(ary)
	fmt.Println(slice)
}
