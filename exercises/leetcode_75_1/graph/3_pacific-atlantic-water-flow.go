/*
	Problem: 417. Pacific Atlantic Water Flow
	Reference: https://leetcode.com/problems/pacific-atlantic-water-flow
    Complexity: O((m*n)*4^n)
*/

type Conclusion struct {
    Atlantic bool
    Pacific bool
}

func (c *Conclusion) Both() bool {
    return c.Atlantic && c.Pacific
}

type Pair struct {
    I, J int
}

type Node struct {
    Pair Pair
    Conclusion Conclusion
    Visited bool
}

func (n *Node) I() int {
    return n.Pair.I
}

func (n *Node) J() int {
    return n.Pair.J
}

func (n *Node) LocatedAtPacific() bool {
    return n.I() == 0 || n.J() == 0
}

func (n *Node) LocatedAtAtlantic(heights [][]int) bool {
    return n.I() == len(heights) - 1 || n.J() == len(heights[n.I()]) - 1
}

func (n *Node) LocatedAtBoth(heights [][]int) bool {
    return n.LocatedAtAtlantic(heights) && n.LocatedAtPacific()
}

func (n *Node) Conclude(conclusions []Conclusion, heights [][]int) Conclusion {
    if n.LocatedAtBoth(heights) {
        return Conclusion{true,true}
    }
    for _, c := range conclusions {
        if c.Both() {
            return Conclusion{true,true}
        }
    }
    atlantic, pacific := n.LocatedAtAtlantic(heights), n.LocatedAtPacific()
    for _, c := range conclusions {
        if c.Atlantic {
            atlantic = true
        }
        if c.Pacific {
            pacific = true
        }
    }
    return Conclusion{atlantic,pacific}
}

func fetchNode(nodes map[Pair]*Node, pair Pair) *Node {
    if _, ok := nodes[pair]; !ok {
        nodes[pair] = &Node{pair, Conclusion{}, false}
    }
    return nodes[pair]
}

func dfs(heights [][]int, nodes map[Pair]*Node, curr, prev Pair) *Node {
    if curr.I < 0 || curr.I >= len(heights) || curr.J < 0 || curr.J >= len(heights[curr.I]) {
        return &Node{}
    }
    if heights[curr.I][curr.J] > heights[prev.I][prev.J] {
        return &Node{}
    }
    node := fetchNode(nodes, curr)
    if node.Visited || node.Conclusion.Both() {
        return node
    }

    node.Visited = true

    l := dfs(heights, nodes, Pair{curr.I-1, curr.J}, curr)
    r := dfs(heights, nodes, Pair{curr.I+1, curr.J}, curr)
    u := dfs(heights, nodes, Pair{curr.I, curr.J-1}, curr)
    d := dfs(heights, nodes, Pair{curr.I, curr.J+1}, curr)

    conclusions := []Conclusion{l.Conclusion, r.Conclusion, u.Conclusion, d.Conclusion}
    conclusion := node.Conclude(conclusions, heights)
    
    node.Conclusion = conclusion
    node.Visited = false

    return node
}

func pacificAtlantic(heights [][]int) [][]int {
    ans := [][]int{}
    nodes := make(map[Pair]*Node)
    for i := range heights {
        for j := range heights[i] {
            pair := Pair{i,j}
            node := dfs(heights, nodes, pair, pair)
            if node.Conclusion.Both() {
                ans = append(ans, []int{i,j})
            }
            node.Visited = true
        }
    }
    return ans
}