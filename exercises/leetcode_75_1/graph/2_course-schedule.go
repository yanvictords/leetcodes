/*
	Problem: 207. Course Schedule
	Reference: https://leetcode.com/problems/course-schedule/
	Results: 321ms, Beats 5.03%
	Complexity: O(V+E)
*/

const (
    PROCESSING = 1
    VISITED = 2
)

func dfsVisited(node int, graph map[int][]int, visited map[int]int) bool {
    if visited[node] == VISITED {
        return false
    }
    if visited[node] == PROCESSING {
        return true
    }
    visited[node] = PROCESSING
    for _, adj := range graph[node] {
        if dfsVisited(adj, graph, visited) {
            return true
        }
    }
    visited[node] = VISITED
    return false
}

func hasCycle(graph map[int][]int) bool {
    for root, _ := range graph {
        visited := make(map[int]int)
        if dfsVisited(root, graph, visited) {
            return true
        }
    }
    return false
}

func buildGraph(prerequisites [][]int) map[int][]int {
    mapper := make(map[int][]int)
    for _, p := range prerequisites {
        mapper[p[1]] = append(mapper[p[1]], p[0])
    }
    return mapper
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := buildGraph(prerequisites)
    return !hasCycle(graph)
}