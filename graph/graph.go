package main

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertext Structure
type Vertex struct {
	key int
	adjecent []*Vertex
}

/// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k })
	}
}


// Add Edge
func (g *Graph) AddEdge(from, to int) {
	// get Vertex
	fromVertex  :=  g.getVertex(from)
	toVertex  :=  g.getVertex(to)
	// check errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalide edge (%v->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjecent, toVertex.key) {
			err := fmt.Errorf("existing edge (%v->%v)", from, to)
			fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjecent = append(fromVertex.adjecent, toVertex)
	}
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

// Print will print the adjuscent lists for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, e := range v.adjecent {
			fmt.Printf(" %d ", e.key)
		} 
	}
	fmt.Println()
}

func main() {
	test := &Graph{}
	for i := 0 ; i < 5 ; i++ {
		test.AddVertex(i)	
	}
	test.AddEdge(1, 2)
	test.AddEdge(1, 3)
	test.AddEdge(4, 1)
	test.AddEdge(3, 4)
	test.AddEdge(3, 5)
	test.AddEdge(2, 3)
	test.AddEdge(2, 5)
	test.AddEdge(2, 1)
	test.Print()
}