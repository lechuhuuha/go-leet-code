package main

import (
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
func insertTreeNode(currentNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < currentNode.key {
		if currentNode.leftNode == nil {
			currentNode.leftNode = newTreeNode
		} else {
			insertTreeNode(currentNode.leftNode, newTreeNode)
		}
	} else {
		if currentNode.rightNode == nil {
			currentNode.rightNode = newTreeNode
		} else {
			insertTreeNode(currentNode.rightNode, newTreeNode)
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

func (tree *BinarySearchTree) PreOrderTree(callback func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTree(tree.rootNode, callback)
}

func preOrderTree(treeNode *TreeNode, callback func(int)) {
	if treeNode == nil {
		return
	}
	callback(treeNode.value)
	preOrderTree(treeNode.leftNode, callback)
	preOrderTree(treeNode.rightNode, callback)
}

// MaxNode method
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		//nil instead of 0
		return nil
	}
	// travese until there is no right node left(biggest node) or tree had 1 element
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

type TreeSet struct {
	bst *BinarySearchTree
}

func (t *TreeSet) Insert(treeNodes ...TreeNode) {
	for _, treeNode := range treeNodes {
		t.bst.InsertElement(treeNode.key, treeNode.value)
	}
}

func (t *TreeSet) Delete(treeNodes ...TreeNode) {
	for _, treeNode := range treeNodes {
		t.bst.RemoveNode(treeNode.key)
	}
}

func (t *TreeSet) Search(treeNodes ...TreeNode) bool {
	for _, treeNode := range treeNodes {
		if exists := t.bst.SearchNode(treeNode.value); !exists {
			return false
		}
	}
	return true
}

func (t *TreeSet) String() {
	t.bst.String()
}

func main() {
	var treeset *TreeSet = &TreeSet{}
	treeset.bst = &BinarySearchTree{}
	var node1 TreeNode = TreeNode{8, 8, nil, nil}
	var node2 TreeNode = TreeNode{3, 3, nil, nil}
	var node3 TreeNode = TreeNode{10, 10, nil, nil}
	var node4 TreeNode = TreeNode{1, 1, nil, nil}
	var node5 TreeNode = TreeNode{6, 6, nil, nil}
	treeset.Insert(node1, node2, node3, node4, node5)
	treeset.String()
}
