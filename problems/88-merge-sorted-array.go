package problems

import (
	"fmt"
	"os"
)

func Merge(nums1 []int, m int, nums2 []int, n int) []int { // original does not have a return and replaces nums1 in place.
	if n == 0 {
		return nums1
	}
	fmt.Fprintln(os.Stdout, nums1, len(nums1))
	fmt.Fprintln(os.Stdout, "========")

	// ~ two pointers problem.
	var p1 int
	if m > 0 {
		p1 = m - 1
	} else {
		p1 = 0
	}
	p2 := n - 1

	lei := m + n - 1 // last element index
	fmt.Fprintln(os.Stdout, lei)

	for p2 >= 0 && lei >= 0 {
		fmt.Fprintln(os.Stdout, "p1,p2,lei", p1, p2, lei)
		fmt.Fprintln(os.Stdout, nums1[p1], nums2[p2])

		fmt.Fprintln(os.Stdout, "----")
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[lei], nums1[p1] = nums1[p1], nums1[lei]
			if p1 > 0 {
				p1--
			}
		} else {
			nums1[lei] = nums2[p2]
			p2--
		}
		lei--
		fmt.Fprintln(os.Stdout, "nums1", nums1)

	}

	fmt.Fprintln(os.Stdout, nums1, len(nums1))
	return nums1

}
