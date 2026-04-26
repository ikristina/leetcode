package problems

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	result := make([]int, len(nums1))

	nge := map[int]int{}
	stack := []int{}

	for _, n := range nums2 {
		for len(stack) > 0 && n > stack[len(stack)-1] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nge[last] = n
		}
		stack = append(stack, n)
	}

	for i, n := range nums1 {
		if _, ok := nge[n]; !ok {
			result[i] = -1
		} else {
			result[i] = nge[n]
		}

	}

	return result

}
