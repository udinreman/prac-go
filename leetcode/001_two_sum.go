package leetcode

import "fmt"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
// UPDATE (2016/2/13):
// The return format had been changed to zero-based indices. Please read the above updated description carefully.

// Subscribe to see which companies asked this question

// TwoSum solution
func TwoSum(nums []int, target int) []int {
	return twoSum2(nums, target)
}

// twoSum1 using two for loop, o(n^2) o(1)
func twoSum1(nums []int, target int) []int {
	for i, n := range nums {
		for j, m := range nums {
			if i != j && m+n == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func twoSum2(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, n := range nums {
		fmt.Println(i, n, indexMap[n])
		if val, ok := indexMap[n]; ok {
			return []int{val, i}
		}
		indexMap[target-n] = i
	}
	return []int{-1, -1}
}
