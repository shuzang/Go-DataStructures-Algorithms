package graph

import (
	"container/heap"
	"container/list"
	"fmt"
)

const MAX_INT = 999999999

type Graph struct {
	VNum      int     // the number of Vertices
	ENum      int     // the numver of Edges
	AdjMatrix [][]int // adjacency matrix
}

type Edge struct {
	from, to int
	weight   int
}

type Heap []Edge

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}
func (h *Heap) Pop() interface{} {
	old := h
	n := len(*old)
	x := (*old)[n-1]
	*h = (*old)[0 : n-1]
	return x
}

func CreateUndirectedGraph() *Graph {
	g := &Graph{}
	g.VNum, g.ENum = 6, 6
	for i := 0; i < 7; i++ {
		//为了便于操作和理解，从下标为1开始
		g.AdjMatrix = append(g.AdjMatrix, make([]int, 7))
	}
	g.AdjMatrix[1][2], g.AdjMatrix[1][3], g.AdjMatrix[1][4] = 1, 1, 1
	g.AdjMatrix[2][1], g.AdjMatrix[2][5] = 1, 1
	g.AdjMatrix[3][1], g.AdjMatrix[3][5] = 1, 1
	g.AdjMatrix[4][1] = 1
	g.AdjMatrix[5][2], g.AdjMatrix[5][3], g.AdjMatrix[5][6] = 1, 1, 1
	g.AdjMatrix[6][5] = 1
	return g
}

// 初始化一个图，顶点和边的数量、权值都预设好
func CreateDirectedGraph() *Graph {
	g := &Graph{}
	g.VNum, g.ENum = 7, 12
	for i := 0; i < 8; i++ {
		//为了便于操作和理解，从下标为1开始
		g.AdjMatrix = append(g.AdjMatrix, make([]int, 8))
	}
	g.AdjMatrix[1][2], g.AdjMatrix[1][4] = 2, 1
	g.AdjMatrix[2][4], g.AdjMatrix[2][5] = 3, 10
	g.AdjMatrix[3][1], g.AdjMatrix[3][6] = 4, 5
	g.AdjMatrix[4][3], g.AdjMatrix[4][5], g.AdjMatrix[4][6], g.AdjMatrix[4][7] = 2, 2, 8, 4
	g.AdjMatrix[5][7] = 6
	g.AdjMatrix[7][6] = 1
	return g
}

func CreateAOV() *Graph {
	g := &Graph{}
	g.VNum, g.ENum = 5, 7
	for i := 0; i < 6; i++ {
		//为了便于操作和理解，从下标为1开始
		g.AdjMatrix = append(g.AdjMatrix, make([]int, 6))
	}
	g.AdjMatrix[1][2], g.AdjMatrix[1][3] = 1, 1
	g.AdjMatrix[2][4], g.AdjMatrix[2][5] = 1, 1
	g.AdjMatrix[3][4], g.AdjMatrix[3][5] = 1, 1
	g.AdjMatrix[4][5] = 1
	return g
}

func CreateAOE() *Graph {
	g := &Graph{}
	g.VNum, g.ENum = 9, 11
	for i := 0; i < 10; i++ {
		//为了便于操作和理解，从下标为1开始
		g.AdjMatrix = append(g.AdjMatrix, make([]int, 10))
	}
	for i := 1; i < 10; i++ {
		for j := 0; j < 10; j++ {
			g.AdjMatrix[i][j] = -1
		}
	}
	g.AdjMatrix[1][2], g.AdjMatrix[1][3], g.AdjMatrix[1][4] = 6, 4, 5
	g.AdjMatrix[2][5] = 1
	g.AdjMatrix[3][5] = 1
	g.AdjMatrix[4][6] = 1
	g.AdjMatrix[5][6], g.AdjMatrix[5][7], g.AdjMatrix[5][8] = 0, 5, 3
	g.AdjMatrix[6][8] = 4
	g.AdjMatrix[7][9] = 2
	g.AdjMatrix[8][9] = 4
	return g
}

func CreateUndirectedWeightedGraph() *Graph {
	g := &Graph{}
	g.VNum, g.ENum = 5, 8
	for i := 0; i < 6; i++ {
		//为了便于操作和理解，从下标为1开始
		g.AdjMatrix = append(g.AdjMatrix, make([]int, 6))
	}
	g.AdjMatrix[1][2], g.AdjMatrix[1][4], g.AdjMatrix[1][5] = 5, 2, 3
	g.AdjMatrix[2][1], g.AdjMatrix[2][3], g.AdjMatrix[2][5] = 5, 1, 3
	g.AdjMatrix[3][2], g.AdjMatrix[3][4], g.AdjMatrix[3][5] = 1, 7, 3
	g.AdjMatrix[4][1], g.AdjMatrix[4][3], g.AdjMatrix[4][5] = 2, 7, 6
	g.AdjMatrix[5][1], g.AdjMatrix[5][2], g.AdjMatrix[5][3], g.AdjMatrix[5][4] = 3, 3, 3, 6
	return g
}

func PrintGraph(g *Graph) {
	for k, v := range (*g).AdjMatrix {
		if k == 0 {
			continue
		}
		fmt.Printf("V%d", k)
		fmt.Println(v)
	}
}

func BreadthFirstSearch(g *Graph, vertex int) []int {
	result := []int{}
	g.AdjMatrix[0][vertex] = 1
	queue := list.New()
	queue.PushBack(vertex)
	for queue.Len() != 0 {
		vertex := queue.Remove(queue.Front()).(int)
		result = append(result, vertex)
		for k, v := range g.AdjMatrix[vertex] {
			if v > 0 && g.AdjMatrix[0][k] != 1 {
				g.AdjMatrix[0][k] = 1
				queue.PushBack(k)
			}
		}
	}
	return result
}

func DepthFirstSearch(g *Graph, vertex int, result []int) []int {
	g.AdjMatrix[0][vertex] = 1
	result = append(result, vertex)
	for k, v := range g.AdjMatrix[vertex] {
		if v > 0 && g.AdjMatrix[0][k] != 1 {
			result = DepthFirstSearch(g, k, result)
		}
	}
	return result
}

func DijkstraShortestPath(g *Graph, vertex int) {
	count := 1                          // 已收录的顶点数目，用于控制循环
	find := make([]bool, g.VNum+1)      //标记已访问过的结点
	prevVertex := make([]int, g.VNum+1) //当前节点的前驱结点
	distance := make([]int, g.VNum+1)   //当前结点的最短路径

	//初始化
	for i := 1; i <= g.VNum; i++ {
		if g.AdjMatrix[vertex][i] > 0 {
			distance[i] = g.AdjMatrix[vertex][i]
			prevVertex[i] = vertex
		} else {
			distance[i] = MAX_INT
			prevVertex[i] = -1
		}
	}

	distance[vertex] = 0
	find[vertex] = true
	v, d := vertex, 0 //用来迭代顶点的变量和初始距离

	for count < g.VNum {
		d = MAX_INT
		for i := 1; i <= g.VNum; i++ {
			if !find[i] && distance[i] < d {
				d = distance[i]
				v = i
			}
		}
		find[v] = true
		count++
		for k, t := range g.AdjMatrix[v] {
			if !find[k] && t > 0 && distance[v]+t < distance[k] {
				distance[k] = distance[v] + t
				prevVertex[k] = v
			}
		}

	}
	fmt.Println(distance[1:])
}

func FloydShortestPath(g *Graph, vertex int) {
	var D [][]int
	var path [][]int
	var i, j, k int
	for i := 0; i < g.VNum; i++ {
		D = append(D, make([]int, g.VNum))
		path = append(path, make([]int, g.VNum))
	}
	for i = 1; i <= g.VNum; i++ {
		for j = 1; j < g.VNum; j++ {
			D[i][j] = g.AdjMatrix[i][j]
			path[i][j] = -1
		}
	}
	for k = 1; k <= g.VNum; k++ {
		for i = 0; i < g.VNum; i++ {
			if D[i][k]+D[k][i] < D[i][j] {
				D[i][j] = D[i][k] + D[k][j]
				path[i][j] = k
			}
		}
	}
}

func TopologicalSort(g *Graph) ([]int, []int) {
	result := make([]int, 1) //拓扑排序的结果数组
	ve := make([]int, g.VNum+1)
	count := 0 //判断图中是否有环

	//计算各结点的入度并存储
	indegree := make([]int, g.VNum+1)
	for i := 1; i <= g.VNum; i++ {
		for j := 1; j <= g.VNum; j++ {
			if g.AdjMatrix[i][j] > 0 {
				indegree[j]++
			}
		}
	}

	queue := list.New()

	//入度为0的结点入队
	for i := 1; i <= g.VNum; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	for queue.Len() != 0 {
		vertex := queue.Remove(queue.Front()).(int)
		result = append(result, vertex)
		count++
		for k, v := range g.AdjMatrix[vertex] {
			if v > 0 {
				indegree[k]--
				if indegree[k] == 0 {
					queue.PushBack(k)
				}
				if ve[vertex]+v > ve[k] {
					ve[k] = ve[vertex] + v
				}
			}
			if v == 0 {
				if ve[vertex]+v > ve[k] {
					ve[k] = ve[vertex] + v
				}
			}
		}
	}
	if count != g.VNum {
		fmt.Println("This is a DAG!")
		return nil, nil
	}
	return result, ve
}

func CriticalPath(g *Graph) (int, []int) {
	path, ve := TopologicalSort(g)
	if len(path) == 1 || path == nil {
		return 0, nil
	}
	vl := make([]int, len(path))
	for i := 1; i < len(vl); i++ {
		vl[i] = MAX_INT
	}
	vl[len(ve)-1] = ve[len(ve)-1]

	for i := len(vl) - 2; i > 0; i-- {
		for k, v := range g.AdjMatrix[i] {
			if v >= 0 && vl[k]-v < vl[i] {
				vl[i] = vl[k] - v
			}
		}
	}
	result := []int{}
	for i := 1; i < g.VNum+1; i++ {
		for j := 1; j < g.VNum+1; j++ {
			if g.AdjMatrix[i][j] >= 0 {
				if ve[i] == vl[j]-g.AdjMatrix[i][j] {
					result = append(result, i)
				}
			}
		}
	}
	result = append(result, path[len(path)-1])
	return ve[len(ve)-1], result

}

func usetFind(x int, uset []int) int {
	for x != uset[x] {
		x = uset[x]
	}
	return x
}

func KruskalMiniSpanTree(g *Graph) (int, []int) {
	var total int
	result := []Edge{}
	h := &Heap{}
	heap.Init(h)

	for i := 1; i < g.VNum+1; i++ {
		for j := i; j < g.VNum+1; j++ {
			if g.AdjMatrix[i][j] > 0 {
				heap.Push(h, Edge{i, j, g.AdjMatrix[i][j]})
			}
		}
	}

	uset := make([]int, g.VNum+1) //用数组表示并查集
	for i := 1; i < len(uset); i++ {
		uset[i] = i
	}

	for h.Len() != 0 {
		e := heap.Pop(h).(Edge)
		if usetFind(e.from, uset) != usetFind(e.to, uset) {
			result = append(result, e)
			uset[uset[e.to]] = uset[e.from]
			total += e.weight
		}
	}
	return total, uset
}

func PrimMiniSpanTree(g *Graph, start int) (int, []int) {
	total := 0
	parent := make([]int, g.VNum+1)
	dist := make([]int, g.VNum+1)

	parent[start] = -1

	for i := 1; i < len(dist); i++ {
		if i == start {
			continue
		}
		if g.AdjMatrix[start][i] > 0 {
			dist[i] = g.AdjMatrix[start][i]
		} else {
			dist[i] = MAX_INT
		}
	}

	count := 1
	vertex, mini := start, MAX_INT

	for count < g.VNum {
		mini = MAX_INT
		for i := 1; i < len(dist); i++ {
			if dist[i] != 0 && dist[i] < mini {
				vertex, mini = i, dist[i]
			}
		}

		total += dist[vertex]
		dist[vertex] = 0
		count++

		for k, t := range g.AdjMatrix[vertex] {
			if dist[k] != 0 && t > 0 && t < dist[k] {
				dist[k] = t
				parent[k] = vertex
			}
		}
	}
	return total, parent
}
