package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}

	tree.Insert(52)
	fmt.Println(tree)

	tree.Insert(310)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(203)
	tree.Insert(150)
	tree.Insert(24)
	tree.Insert(276)
	tree.Insert(7)
	tree.Insert(88)

	fmt.Println(tree)

	fmt.Println(tree.Search(211))
	fmt.Println(tree.Search(19))
	fmt.Println(tree.Search(276))
}
