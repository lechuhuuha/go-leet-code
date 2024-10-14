package main

import (
	"bytes"
	"fmt"
	"sync"
)

// TreeNode class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

// InsertElement method
func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	treeNode := &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

// insertTreeNode function
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

// inOrderTraverseTree method
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// MaxNode method
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		//nil instead of 0
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// SearchNode method
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

// searchNode method
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

// RemoveNode method
func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, key)
}

// Helper function to find the smallest node in the right subtree.
func findMinNode(treeNode *TreeNode) *TreeNode {
	for treeNode != nil && treeNode.leftNode != nil {
		treeNode = treeNode.leftNode
	}
	return treeNode
}

// Improved removeNode method
func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	// Traverse the tree to find the node to remove
	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}
	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	// Node found: handle removal based on the number of children

	// Case 1: Node has no children (leaf node)
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		return nil
	}

	// Case 2: Node has only one child
	if treeNode.leftNode == nil {
		return treeNode.rightNode
	}
	if treeNode.rightNode == nil {
		return treeNode.leftNode
	}

	// Case 3: Node has two children
	// Find the smallest value in the right subtree (in-order successor)
	minNode := findMinNode(treeNode.rightNode)
	treeNode.key, treeNode.value = minNode.key, minNode.value

	// Remove the in-order successor
	treeNode.rightNode = removeNode(treeNode.rightNode, minNode.key)
	return treeNode
}

// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(tree.rootNode, 0)
	fmt.Println("------------------------------------------------")
}

// stringify method
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---[ "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

// main method
func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.String()

	tr := New(NewNode(Int(5)))
	tr.Insert(NewNode(Int(3)))
	tr.Insert(NewNode(Int(17)))
	tr.Insert(NewNode(Int(7)))
	tr.Insert(NewNode(Int(1)))
	/*
	       5
	      / \
	     3   17
	    /    /
	   1    7
	*/

	fmt.Println("Min:", tr.Min().Key)             // Min: 1
	fmt.Println("Max:", tr.Max().Key)             // Max: 17
	fmt.Println("Search:", tr.Search(Int(7)).Key) // Search: 7

	buf1 := new(bytes.Buffer)
	ch1 := make(chan string)
	go tr.PreOrder(ch1) // root, left, right
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		buf1.WriteString(v)
		buf1.WriteString(" ")
	}
	fmt.Println("PreOrder:", buf1.String()) // PreOrder: 5 3 1 17 7

	buf2 := new(bytes.Buffer)
	ch2 := make(chan string)
	go tr.InOrder(ch2) // left, root, right
	for {
		v, ok := <-ch2
		if !ok {
			break
		}
		buf2.WriteString(v)
		buf2.WriteString(" ")
	}
	fmt.Println("InOrder:", buf2.String()) // 1 3 7 17 5

	buf3 := new(bytes.Buffer)
	ch3 := make(chan string)
	go tr.PostOrder(ch3) // left, right, root
	for {
		v, ok := <-ch3
		if !ok {
			break
		}
		buf3.WriteString(v)
		buf3.WriteString(" ")
	}
	fmt.Println("PostOrder:", buf3.String()) // 1 3 7 17 5

	buf4 := new(bytes.Buffer)
	nodes4 := tr.LevelOrder() // from top-level
	for _, v := range nodes4 {
		buf4.WriteString(fmt.Sprintf("%v ", v.Key))
	}
	fmt.Println("LevelOrder:", buf4.String()) // 5 3 17 1 7

	tr2 := New(NewNode(Int(5)))
	tr2.Insert(NewNode(Int(3)))
	tr2.Insert(NewNode(Int(17)))
	tr2.Insert(NewNode(Int(7)))
	tr2.Insert(NewNode(Int(1)))

	fmt.Println("ComparePreOrder:", ComparePreOrder(tr, tr2))   // true
	fmt.Println("CompareInOrder:", CompareInOrder(tr, tr2))     // true
	fmt.Println("ComparePostOrder:", ComparePostOrder(tr, tr2)) // true
}
