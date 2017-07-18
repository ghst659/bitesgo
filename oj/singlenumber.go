package oj

// Given an array of integers, every element appears twice except for one.
// Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity.
// Could you implement it without using extra memory?

func singleNumber(nums []int) int {
	var result int
	record := make(map[int]bool)
	for _, k := range nums {
		record[k] = !record[k]
	}
	for k, _ := range record {
		if record[k] {
			result = k
			break
		}
	}
	return result
}
