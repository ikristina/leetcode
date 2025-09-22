package problems

import "fmt"

// Given an array nums of size n, return the majority element.
//The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {
	hm := map[int]int{}
	for _, el := range nums {
		if _, ok := hm[el]; !ok {
			hm[el] = 1
		} else {
			hm[el]++
		}
	}
	me := 0
	el := 0
	for k, v := range hm {
		fmt.Println(k, v)
		if v > me {
			me = v
			el = k
		}
	}

	return el
}
