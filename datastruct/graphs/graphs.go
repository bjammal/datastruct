package graphs

//Graph represents a graph of nodes and edges
type Graph struct {
	Nodes map[int][]int
}

//Gnode represents graph node
type Gnode struct {
	Val       int
	adjacents []*Gnode
}

//AddEdge inserts an edge from n
func (g *Graph) AddEdge(u, v int) bool {
	if n, ok := g.nodes[u]; ok {
		n = append(n, v)
		return true
	}
	return false
}
