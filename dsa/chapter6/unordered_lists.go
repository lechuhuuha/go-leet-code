package main

import "fmt"

type UnorderedList struct {
	headNode *Node
}

func (u *UnorderedList) AddToHead(p int) {
	node := Node{
		property: rune(p),
	}
	if u.headNode != nil {
		node.nextNode = u.headNode
	}
	u.headNode = &node
}

func (u *UnorderedList) IterateList() {
	for node := u.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}
