/*
	Problem: 295. Find Median from Data Stream
	Reference: https://leetcode.com/problems/find-median-from-data-stream/
    Complexity: O(N log N)
*/

type Heap struct {
    Left []int
    Right []int
}

func NewHeap() *Heap {
    return &Heap{
        Left: []int{},
        Right: []int{},
    }
}

func (h *Heap) Swap(arr []int, i,j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func (h *Heap) Len(arr []int) int {
    return len(arr)
}

func (h *Heap) Even() bool {
    return h.Len(h.Left) == h.Len(h.Right)
}

func (h *Heap) Median() float64 {
    return float64(h.Peek(h.Left) + h.Peek(h.Right))/2
}

func (h *Heap) Mid() float64 {
    if h.Len(h.Left) > h.Len(h.Right) {
        return float64(h.Peek(h.Left))
    }
    return float64(h.Peek(h.Right))
}

func (h *Heap) Peek(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    return arr[0]
}

func (h *Heap) Push(val int) {
    if val > h.Peek(h.Right) {
        rightVal := val
        if h.Len(h.Left) < h.Len(h.Right) {
            rightVal = h.pop(&h.Right, h.siftDownRight)
            h.push(&h.Left, min(val,rightVal), h.siftUpLeft)
        }
        h.push(&h.Right, max(val,rightVal), h.siftUpRight)
        return
    }
    leftVal := val
    if h.Len(h.Left) > h.Len(h.Right) {
        leftVal = h.pop(&h.Left, h.siftDownLeft)
        h.push(&h.Right, max(val, leftVal), h.siftUpRight)
    }
    h.push(&h.Left, min(val, leftVal), h.siftUpLeft)
    return
}

func (h *Heap) push(arr *[]int, val int, siftUp func(idx int)) {
    *arr = append(*arr, val)
    siftUp(h.Len(*arr)-1)
}

func (h *Heap) pop(arr *[]int, siftDown func(idx int)) int {
    val := (*arr)[0]
    (*arr)[0] = (*arr)[len(*arr)-1]
    *arr = (*arr)[:len(*arr)-1]
    siftDown(0)
    return val
}

func (h *Heap) siftUpLeft(idx int) {
    for idx > 0 {
        parent := (idx-1) / 2
        if h.Left[idx] < h.Left[parent] {
            break
        }
        h.Swap(h.Left, idx, parent)
        idx = parent
    }
}

func (h *Heap) siftUpRight(idx int) {
    for idx > 0 {
        parent := (idx-1) / 2
        if h.Right[idx] > h.Right[parent] {
            break
        }
        h.Swap(h.Right, idx, parent)
        idx = parent
    } 
}

func (h *Heap) siftDownLeft(idx int) {
    for idx < len(h.Left) {
        leftChildren := (idx*2) + 1
        rightChildren := (idx*2) + 2
        children := leftChildren
        if rightChildren < len(h.Left) && h.Left[rightChildren] > h.Left[leftChildren] {
            children = rightChildren
        }
        if children >= len(h.Left) || h.Left[idx] > h.Left[children] {
            break
        }
        h.Swap(h.Left, idx, children)
        idx = children
    }   
}

func (h *Heap) siftDownRight(idx int) {
    for idx < len(h.Right) {
        leftChildren := (idx*2) + 1
        rightChildren := (idx*2) + 2
        children := leftChildren
        if rightChildren < len(h.Right) && h.Right[rightChildren] < h.Right[leftChildren] {
            children = rightChildren
        }
        if children >= len(h.Right) || h.Right[idx] < h.Right[children] {
            break
        }
        h.Swap(h.Right, idx, children)
        idx = children
    }
}


type MedianFinder struct {
    heap *Heap
}

func Constructor() MedianFinder {
    return MedianFinder{
        heap: NewHeap(),
    }
}

func (this *MedianFinder) AddNum(num int) {
    this.heap.Push(num)
}


func (this *MedianFinder) FindMedian() float64 {
   if this.heap.Even() {
        return this.heap.Median()
   }
    return this.heap.Mid()
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */