package problems

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// binary search on the smaller array to keep j valid
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := len(nums1)
	n := len(nums2)

	// binary search over possible partition position in nums1
	low, high := 0, m

	for low <= high {
		// i = how many elements we take from nums1's left side
		i := (low + high)/2
		// j = how many elements we take from nums2's left side (forced by i)
		// +1 ensures the left side is larger on odd total, so max(left) is the median
		// i + j = half of all the elements in both arrays.
	 	j := (m + n + 1)/2 - i

		// defaults: -inf for left edges (no left element = always valid)
        //           +inf for right edges (no right element = always valid)

        left1, left2 := math.MinInt, math.MinInt
        right1, right2 := math.MaxInt, math.MaxInt

        if i > 0 {
       		left1 = nums1[i - 1] // last element of nums1's left side
        }
        if i < m {
        	right1 = nums1[i] // first element of nums1's right side
        }
        if j > 0 {
        	left2 = nums2[j - 1] // last element of nums2's left side
        }
        if j < n {
        	right2 = nums2[j] // first element of nums2's right side
        }

        if left1 > right2 {
        	// took too many from nums1, move partition to the left
        	high = i - 1
        } else if left2 > right1 {
        	// took too few from nums1, move partition right
        	low = i + 1
        } else {
        	// correct partition: all left elements <= all right elements
        	if (m + n) % 2 == 0 {
        		// even total: average the two middle elements
        		return (float64(min(right1, right2)) + float64(max(left1, left2)))/2
        	} else {
        		return float64(max(left1, left2))
        	}
        }
	}
	return 0
}
