/*
	Problem: 42. Trapping Rain Water
	Reference: https://leetcode.com/problems/insert-delete-getrandom-o1
	Results: 154ms, Beats 57.59%
	Complexity: O(1)
*/

type RandomizedSet struct {
	Idx map[int]int
	Set map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Idx: make(map[int]int),
		Set: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Idx[val]; ok {
		return false
	}
	idx := len(this.Idx)
	this.Idx[val] = idx
	this.Set[idx] = val
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Idx[val]; !ok {
		return false
	}
	idx := this.Idx[val]
	delete(this.Idx, val)
	delete(this.Set, idx)
	if len(this.Set) == idx {
		return true
	}
	// fills the empty space with the last val
	lastIdx := len(this.Idx)
	lastVal := this.Set[lastIdx]
	this.Idx[lastVal] = idx
	this.Set[idx] = lastVal
	delete(this.Set, lastIdx)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	random := rand.IntN(len(this.Set))
	return this.Set[random]
}

/**
 * Your RandomizedIdx object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */