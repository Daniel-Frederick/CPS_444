
package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

// Insert inserts a value into the BST
func (node *Node) Insert(val int) *Node {
	if node == nil {
		return &Node{value: val}
	}

	if val < node.value {
		node.left = node.left.Insert(val)
	} else if val > node.value {
		node.right = node.right.Insert(val)
	}
	// no duplicates
	return node
}

// Find checks if a value exists in the BST
func (node *Node) Find(val int) bool {
	if node == nil {
		return false
	}

	if val == node.value {
		return true
	} else if val < node.value {
		return node.left.Find(val)
	} else {
		return node.right.Find(val)
	}
}

// IsEmpty checks if the tree is empty
func (node *Node) IsEmpty() bool {
	return node == nil
}

// InOrderTraversal returns the values in sorted order
func (node *Node) InOrderTraversal() []int {
	if node == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, node.left.InOrderTraversal()...)
	result = append(result, node.value)
	result = append(result, node.right.InOrderTraversal()...)
	return result
}

func main() {
	var tree *Node

	fmt.Println("Is tree empty:", tree.IsEmpty())

	fmt.Println("Inserting values...")
	tree = tree.Insert(10)
	tree = tree.Insert(5)
	tree = tree.Insert(15)
	tree = tree.Insert(3)
	tree = tree.Insert(8)

	fmt.Println("Is tree empty:", tree.IsEmpty())

	fmt.Println("Find 8:", tree.Find(8))
	fmt.Println("Find 100:", tree.Find(100))

	fmt.Println("InOrder Traversal:", tree.InOrderTraversal())
}

