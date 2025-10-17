package problems

import "math/rand/v2"

type RandomizedSet struct {
	vals []int
	m    map[int]int // {value:index in vals}
}

func Constructor() RandomizedSet {
	rs := RandomizedSet{}
	rs.m = make(map[int]int)
	return rs
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.m[val]; !ok {
		rs.vals = append(rs.vals, val)
		rs.m[val] = len(rs.vals) - 1
		return true
	}
	return false
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.m[val]; ok {
		lastItem := rs.vals[len(rs.vals)-1]
		rs.vals[rs.m[val]] = lastItem
		rs.vals = rs.vals[:len(rs.vals)-1]
		rs.m[lastItem] = rs.m[val]
		delete(rs.m, val)
		return true
	}
	return false
}

func (rs *RandomizedSet) GetRandom() int {
	if len(rs.vals) == 0 {
		return 0
	}
	return rs.vals[rand.IntN(len(rs.vals))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
