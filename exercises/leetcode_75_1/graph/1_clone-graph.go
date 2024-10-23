/*
	Problem: 133. Clone Graph
	Reference: https://leetcode.com/problems/clone-graph/
	Results: 0ms, Beats 100.00%
	Complexity: O(V+E)
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func dfs(node *Node, references map[int]*Node) *Node {
    if node == nil {
        return nil
    }
    if references[node.Val] != nil {
        return references[node.Val]
    }
    clone := &Node{Val: node.Val}
    references[clone.Val] = clone
    for _, n := range node.Neighbors {
        clone.Neighbors = append(clone.Neighbors, dfs(n, references))
    }
    return clone
}

func cloneGraph(node *Node) *Node {
    references := make(map[int]*Node)
    return dfs(node, references)
}