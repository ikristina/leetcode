package problems

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			nums1: []int{},
			nums2: []int{},
			want:  0.0,
		},
	}

	for _, test := range tests {
		if got := findMedianSortedArrays(test.nums1, test.nums2); got != test.want {
			t.Errorf("FindMedianSortedArrays(%v, %v) = %v, want %v", test.nums1, test.nums2, got, test.want)
		}
	}
}
