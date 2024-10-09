package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property

	if node.nextNode != nil {
		node.nextNode = l.headNode
	}

	l.headNode = &node
}

func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (l *LinkedList) LastNode() *Node {
	var (
		node     *Node
		lastNode *Node
	)
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

func (l *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	lastNode := l.LastNode()
	if lastNode == nil {
		return
	}
	lastNode.nextNode = node
}

func (l *LinkedList) NodeWithValue(property int) *Node {
	var (
		nodeWith *Node
	)
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	var (
		node     = &Node{}
		nodeWith *Node
	)

	node.property = property

	nodeWith = l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()
}
