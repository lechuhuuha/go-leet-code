package main

import "fmt"

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
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

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode == nil && node.nextNode == nil {
			continue
		}

		if node.previousNode.property != firstProperty &&
			node.nextNode.property != secondProperty {
			continue
		}

		nodeWith = node
		break
	}
	return nodeWith
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

func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	var nodeWith *Node

	node.property = property
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	node := linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
