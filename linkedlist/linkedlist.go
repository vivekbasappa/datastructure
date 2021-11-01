// This is example code written based on Junmin Lee video tutorial
// All credit goes to her. I am trying just to understand datastructures and
// implemention in go

package main

import "fmt"

// Node represents single entity in single linked list
type Node struct {
	Data int
	Next *Node
}

func NewNode(value int)  *Node {
	node := new(Node)
	node.Data = value
	return node
}

//LinkedList represents single linked list 
type LinkedList struct {
	Head *Node // points to head of the list
	Length int // total nodes in the list
}

// Prepend add node at the beginning of the list
func (l *LinkedList) Prepend(node *Node) {
	second := l.Head
	l.Head = node
	l.Head.Next = second
	l.Length++
}

func(l *LinkedList) Reverse()  {
	currentNode := l.Head	
	var next *Node
	var previousNode *Node
	
	for currentNode != nil {
		next, currentNode.Next = currentNode.Next, previousNode
		previousNode, currentNode = currentNode, next
	}
	l.Head = previousNode
}

// Delete the node from the linked list
func (l *LinkedList) DeleteByValue(value int)  error {
	if l.Head == nil {
		return fmt.Errorf("error no elements in the linked list")
	} else if l.Head.Data == value {
		l.Head = l.Head.Next
		return nil
	}
	node := l.Head
	for ; node.Next != nil && node.Next.Data != value ; node = node.Next {
	}
	if node.Next == nil {
		return fmt.Errorf("error unable to find value %d", value)
	}
	node.Next = node.Next.Next
	return nil
}

// Print print all the elements in the list
func (l LinkedList) Print() {
	for node := l.Head ; node != nil ; node = node.Next{
		fmt.Printf("%d ", node.Data)
	}
	fmt.Println()
}

// NewLinkedList constructor
func NewLinkedList() *LinkedList {
	linkedList := new(LinkedList)
	return linkedList
}

func main1() {
	linkedList := NewLinkedList()	
	node := NewNode(10)
	linkedList.Prepend(node)
	node = NewNode(20)
	linkedList.Prepend(node)
	node = NewNode(30)
	linkedList.Prepend(node)
	node = NewNode(40)
	linkedList.Prepend(node)
	node = NewNode(50)
	linkedList.Prepend(node)

	/*
	linkedList.Print()
	linkedList.DeleteByValue(30)
	linkedList.Print()
	linkedList.DeleteByValue(50)
	linkedList.Print()
	linkedList.DeleteByValue(10)
	linkedList.Print()
	err := linkedList.DeleteByValue(100)
	if err != nil {
		fmt.Println(err.Error())
	}
	*/
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
	fmt.Println()
}