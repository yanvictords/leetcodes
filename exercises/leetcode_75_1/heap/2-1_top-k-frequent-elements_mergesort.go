/*
	Problem: 347. Top K Frequent Elements
	Reference: https://leetcode.com/problems/top-k-frequent-elements/
    Complexity: O(N log N)
*/

type Number struct {
    Value int
    Frequency int
}

type Set struct {
    Values []*Number
}

func NewSet(mapper map[int]int) *Set {
    values := make([]*Number, len(mapper))
    key := 0
    for value, frequency := range mapper {
        values[key] = &Number{value, frequency}
        key++
    }
    return &Set{Values: values}
}

func (s *Set) Len() int {
    return len(s.Values)
}

func (s *Set) Sort() []*Number {
    merge := func (left []*Number, right []*Number) []*Number {
        merged := make([]*Number, 0)
        l, r := 0, 0
        for l < len(left) && r < len(right) {
            if left[l].Frequency < right[r].Frequency {
                merged = append(merged, left[l])
                l++
                continue
            }
            merged = append(merged, right[r])
            r++
        }

        for ; l < len(left); l++ {
            merged = append(merged, left[l])
        }
        for ; r < len(right); r++ {
            merged = append(merged, right[r])
        }

        return merged
    }
    var mergeSort func (list []*Number) []*Number
    mergeSort = func (list []*Number) []*Number {
        if len(list) < 2 {
            return list
        }
        left := mergeSort(list[:len(list)/2])
        right := mergeSort(list[len(list)/2:])
        return merge(left, right)
    }

    s.Values = mergeSort(s.Values)
    return s.Values
}

func (s *Set) Top(k int) []int {
    ans := make([]int, k)
    for i := 0; i < k; i++ {
        ans[i] = s.Values[len(s.Values)-1-i].Value
    }
    return ans
}

func topKFrequent(nums []int, k int) []int {
    hitMapper := make(map[int]int, len(nums))
    for _, n := range nums {
        hitMapper[n]++
    }
    set := NewSet(hitMapper)
    set.Sort()
    return set.Top(k)
}