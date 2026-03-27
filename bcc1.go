package main
import "fmt"
type Edge struct {
	u, v int
}

var time int

func biconnectedComponents(n int, graph [][]int){
	disc := make([]int, n)
	low := make([]int, n)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	var st []Edge

	var dfs func(u int)
	dfs = func(u int){
		children := 0
		time++
		disc[u] = time
		low[u] = time
		for _, v := range graph[u]{
			if disc[v]==0{
				children++
				parent[v]=u
				st = append(st, Edge{u,v})
				dfs(v)
				low[u] = min(low[u], low[v])
			}
		}
	}
}

//biconnected components
//a graph is biconnected if you can remove any one vertex and the graph is still connected
//a point is an articulation point if removing it breaks the graph
//a biconnected component is a maximal subgraph with no articulation point inside; a "robust cluster"
