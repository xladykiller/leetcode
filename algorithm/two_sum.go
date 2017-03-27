package algorithm

/**
https://leetcode.com/problems/two-sum/
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */
func TwoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i, v := range nums {

		idx, ok := cache[target-v]
		if ok && idx != i {
			return []int{i, idx}
		}

		cache[v] = i

	}
	return nil

}
