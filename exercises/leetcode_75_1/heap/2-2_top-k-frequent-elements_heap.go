/*
	Problem: 347. Top K Frequent Elements
	Reference: https://leetcode.com/problems/top-k-frequent-elements/
    Complexity: O(N log N)
*/

type Number struct {
    Value int
    Frequency int
}

type Heap struct {
    Tree []*Number
}

func NewHeap() *Heap {
    return &Heap {
        Tree: []*Number{},
    }
}

func (h *Heap) Len() int {
    return len(h.Tree)
}

func (h *Heap) Swap(i,j int) {
    h.Tree[i], h.Tree[j] = h.Tree[j], h.Tree[i]
}

func (h *Heap) Push(number *Number) {
    h.Tree = append(h.Tree, number)
    h.siftUp(h.Len()-1)
}

func (h *Heap) Pop() *Number {
    if h.Len() == 0 {
        return nil
    }
    number := h.Tree[0]
    h.Tree[0] = h.Tree[h.Len()-1]
    h.Tree = h.Tree[:h.Len()-1]
    h.siftDown(0)
    return number
}

func (h *Heap) siftUp(idx int) {
    for idx > 0 {
        parent := (idx-1) / 2
        if h.Tree[idx].Frequency < h.Tree[parent].Frequency {
            break
        }
        h.Swap(idx, parent)
        idx = parent
    }
}

func (h *Heap) siftDown(idx int) {
    for idx < h.Len() {
        left := (idx*2) + 1
        right := (idx*2) + 2
        children := left
        if right < h.Len() && h.Tree[left].Frequency < h.Tree[right].Frequency {
            children = right
        }
        if children >= h.Len() || h.Tree[idx].Frequency > h.Tree[children].Frequency {
            break
        }
        h.Swap(idx, children)
        idx = children
    }
}

func topKFrequent(nums []int, k int) []int {
    hitMapper := make(map[int]int, len(nums))
    for _, num := range nums {
        hitMapper[num]++
    }
    heap := NewHeap()
    for value, frequency := range hitMapper {
        number := &Number{value, frequency}
        heap.Push(number)
    }
    ans := make([]int, k)
    for i := 0; i < k; i++ {
        number := heap.Pop()
        ans[i] = number.Value
    }
    return ans
}