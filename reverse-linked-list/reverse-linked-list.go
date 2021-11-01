package main

import "fmt"

// Node structure
type Node struct {
	value int // node data
	next *Node
}

// newNode create new node
func newNode(value int, next *Node)  *Node {
	node :=  new(Node)
	node.value = value
	node.next = next
	return node
}

// LinkedList structure
type LinkedList struct {
	root *Node
}

//Insert in to linked list
func(l *LinkedList) Insert(value int) {
	if l.root == nil {
		l.root = newNode(value, nil)
		return
	}
	node := l.root
	for  ; node.next != nil ; node = node.next {
	}
	node.next = newNode(value, nil)
}

//Print print content of linkedlist values
func(l *LinkedList) Print() {
	fmt.Println()	
	for  node := l.root; node != nil ; node = node.next {
		fmt.Printf(" %d ",  node.value)
	}
	fmt.Println()	
}

//Remove from linked list
func(l *LinkedList) Remove(value int) {
	// error handling if root empty 
	if l.root == nil {
		err := fmt.Errorf("list empty")
		fmt.Println(err.Error())
	}

	// if root contains value
	if l.root.value == value {
		l.root = l.root.next
	}	

	node := l.root
	for ; node.next != nil  && node.next.value != value ; node = node.next {
	}
	if node.next == nil {
		err := fmt.Errorf("value not found")
		fmt.Println(err.Error())
	}
}

func (l *LinkedList) Reverse() {
	currentNode := l.root
	var previousNode *Node = nil
	var nextNode *Node = nil

	for ; currentNode != nil ;  {
		nextNode, currentNode.next = currentNode.next, previousNode
		previousNode, currentNode = currentNode, nextNode	
	}
	l.root = previousNode
}

func main() {
	l := &LinkedList{}
	l.Insert(10)
	l.Insert(20)
	l.Insert(30)
	l.Insert(40)
	l.Print()
	l.Reverse()
	l.Print()
}