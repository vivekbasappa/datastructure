package main

import "fmt"

// Tree
type Node struct {
	Data int
	Right *Node
	Left *Node
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.Left == nil {
			n.Left = &Node{ Data: data, Left:nil , Right: nil}
		} else {
			n.Left.insert(data)
		} 
	} else {
		if n.Right == nil {
			n.Right = &Node{ Data: data, Left:nil , Right: nil}
		} else {
			n.Right.insert(data)
		}
	}
}

type Tree struct {
	root *Node
}

func(t *Tree) Insert(value int) {
		if t.root == nil {
			t.root = &Node{}
			t.root.Data = value
		} else {
			t.root.insert(value)
		}
}

func Inorder(t *Node) {
	if t == nil {
		return
	}
	Inorder(t.Left)
	fmt.Printf(" %d ", t.Data)
	Inorder(t.Right)
}



// BFS 
func BFS(node *Node) []int {
	queue := []*Node{}
	queue = append(queue, node)
	result := []int{}
	res := BFSUtil(queue, result)
	return res
}

// BFS Utility program
func BFSUtil(queue []*Node, result []int) []int {
	if len(queue) == 0 {
		return result
	}
	result = append(result, queue[0].Data)
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	return BFSUtil(queue[1:], result)
}

func main() {
	tree := &Tree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(60)
	Inorder(tree.root)
	fmt.Println()
	result := BFS(tree.root)
	fmt.Println(result)
}