package oj

import (
	"testing"
)

// Given an array of integers, every element appears twice except for one.
// Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity.
// Could you implement it without using extra memory?

func TestSingleNumber(t *testing.T) {
	cases := [][]int{
		[]int{0, 1, 1},
		[]int{0, 2, 1, 2, 0},
		[]int{0, 2, 1, 1, 3, 0, 3},
	}
	for i, nums := range cases {
		val := singleNumber(nums)
		if val != i {
			t.Errorf("incorrect value, got %d, want %d", val, i)
		}
	}
}
