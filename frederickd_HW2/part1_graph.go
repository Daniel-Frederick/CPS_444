package main

import (
	"fmt"
)

type Graph struct {
	vertices int
	matrix   [][]int
	directed bool
}

func NewGraph(vertices int, directed bool) *Graph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	
	return &Graph{
		vertices: vertices,
		matrix:   matrix,
		directed: directed,
	}
}

func (graph *Graph) AddEdge(from, to int) {
	validFrom := from >= 0 && from < graph.vertices
	validTo := to >= 0 && to < graph.vertices
	if validFrom && validTo {
		graph.matrix[from][to] = 1
		if !graph.directed {
			graph.matrix[to][from] = 1
		}
	}
}

func (graph *Graph) RemoveEdge(from, to int) {
	validFrom := from >= 0 && from < graph.vertices
	validTo := to >= 0 && to < graph.vertices
	if validFrom && validTo {
		graph.matrix[from][to] = 0
		if !graph.directed {
			graph.matrix[to][from] = 0
		}
	}
}

func (graph *Graph) HasEdge(from, to int) bool {
	validFrom := from >= 0 && from < graph.vertices
	validTo := to >= 0 && to < graph.vertices
	if validFrom && validTo {
		return graph.matrix[from][to] == 1
	}
	return false
}

func (g *Graph) Print() {
	isDirectedMap := map[bool]string{true: "Directed", false: "Undirected"}
	fmt.Printf("Adjacency Matrix (%s):\n", isDirectedMap[g.directed])
	fmt.Print("   ")
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%2d ", i)
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%2d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// Non-Directed Graph
	undirectedGraph := NewGraph(4, false)
	undirectedGraph.AddEdge(0, 1)
	undirectedGraph.AddEdge(1, 2)
	undirectedGraph.Print()
	
	fmt.Println()
	
	// Directed Graph
	directedGraph := NewGraph(4, true)
	directedGraph.AddEdge(0, 1)
	directedGraph.AddEdge(1, 2)
	directedGraph.Print()
}
