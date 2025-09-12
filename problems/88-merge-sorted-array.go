package problems

func Merge(nums1 []int, m int, nums2 []int, n int) []int { // original does not have a return and replaces nums1 in place.
	if n == 0 {
		return nums1
	}

	// ~ two pointers problem.

	p1 := m - 1
	p2 := n - 1
	lei := m + n - 1 // last element index

	for p2 >= 0 {

		if p1 >= 0 && nums1[p1] >= nums2[p2] {
			nums1[lei] = nums1[p1]
			p1--
		} else {
			nums1[lei] = nums2[p2]
			p2--
		}
		lei--
		if lei < 0 {
			break
		}
	}

	return nums1

}
