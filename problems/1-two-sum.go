package problems

// TwoSum leetcode.com/problems/two-sum
// easy
// #array, #hash table
func TwoSum(nums []int, target int) []int {
	hm := map[int]int{}
	for i, v := range nums {
		diff := target - v
		if _, ok := hm[diff]; !ok {
			hm[v] = i
			continue
		}
		return []int{hm[diff], i}
	}
	return []int{0, 0}
}
