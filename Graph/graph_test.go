package graph

import (
	"fmt"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	PrintGraph(CreateDirectedGraph())
	PrintGraph(CreateUndirectedGraph())
	PrintGraph(CreateAOE())
}

func TestBreadthFirstSearch(t *testing.T) {
	g := CreateUndirectedGraph()
	PrintGraph(g)
	fmt.Println("BFS的顺序为: ", BreadthFirstSearch(g, 1))
}

func TestDepthFirstSearch(t *testing.T) {
	g := CreateUndirectedGraph()
	PrintGraph(g)
	result := []int{}
	fmt.Println("DFS的顺序为: ", DepthFirstSearch(g, 1, result))
}

func TestDijkstraShortestPath(t *testing.T) {
	g := CreateDirectedGraph()
	PrintGraph(g)
	fmt.Printf("各结点的最短路径为:")
	DijkstraShortestPath(g, 1)
}

func TestTopologicalSort(t *testing.T) {
	g := CreateAOV()
	PrintGraph(g)
	result, _ := TopologicalSort(g)
	fmt.Println("拓扑排序的结果为: ", result[1:])
}

func TestCriticalPath(t *testing.T) {
	g := CreateAOE()
	PrintGraph(g)
	total, result := CriticalPath(g)
	fmt.Println("工程总用时为: ", total)
	fmt.Println("关键路径为: ", result)
}

func TestKruskalMiniSpanTree(t *testing.T) {
	g := CreateUndirectedWeightedGraph()
	PrintGraph(g)
	total, uset := KruskalMiniSpanTree(g)
	fmt.Println("总权重为: ", total)
	fmt.Println("并查集为: ", uset)
}

func TestPrimMiniSpanTree(t *testing.T) {
	g := CreateUndirectedWeightedGraph()
	PrintGraph(g)
	total, uset := PrimMiniSpanTree(g, 1)
	fmt.Println("总权重为: ", total)
	fmt.Println("各结点的父节点分别为(-1表示根结点): ", uset[1:])
}
