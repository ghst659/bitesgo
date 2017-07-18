package oj

// Given an array of integers, every element appears twice except for one.
// Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity.
// Could you implement it without using extra memory?

func singleNumber(nums []int) int {
	var result int
	record := make(map[int]int)
	for _, k := range nums {
		if _, seen := record[k]; seen {
			delete(record, k)
		} else {
			record[k] = 1
		}
	}
	for result, _ = range record {
		break
	}
	return result
}
