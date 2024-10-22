package main

type Node struct {
	nextNode *Node
	property rune
}

func CreateLinkedList() *Node {
	headNode := Node{
		property: 'a',
	}

	currentNode := headNode

	for i := 'b'; i < 'z'; i++ {
		currentNode.nextNode = &Node{property: i}
		currentNode = *currentNode.nextNode
	}

	return &headNode
}

func ReverseLinkedList(nodeList *Node) *Node {
	currNode := nodeList
	var topNode *Node

	for {
		if currNode == nil {
			break
		}
		tempNode := currNode.nextNode
		currNode.nextNode = topNode
		topNode = currNode
		currNode = tempNode
	}

	return topNode
}
