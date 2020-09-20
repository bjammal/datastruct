package sorting

//TopologicalSort returns the topological sort of a graph.
//Graph is represented with an adjacency list using a map.
//This simplified implementation assumes a directed acyclic graph.
func TopologicalSort(g map[int][]int) ([]int, error) {
	visited := make(map[int]bool)
	stack := []int{}
	for node := range g {
		visitNode(g, node, visited, &stack)
	}
	return stack, nil
}

func visitNode(graph map[int][]int, node int, visited map[int]bool, s *[]int) {
	if !visited[node] {
		visited[node] = true
		for _, n := range graph[node] {
			visitNode(graph, n, visited, s)
		}
		*s = append(*s, node)
	}
}
