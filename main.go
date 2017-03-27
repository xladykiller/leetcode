package main

import (
	"fmt"
	"github.com/xladykiller/leetcode/algorithm"
)

func main() {

	out := algorithm.TwoSum([]int{3, 3, 4}, 6)
	fmt.Println("two sum:", out)

	result := algorithm.AddTwoNumbers(algorithm.BuildLink([]int{0}), algorithm.BuildLink([]int{7, 3}))
	fmt.Print("add two numbers: ")
	for result != nil {
		fmt.Print(result.Val)
		if result.Next != nil {
			fmt.Print("->")
		} else {
			fmt.Println("")
		}
		result = result.Next
	}

	length := algorithm.LengthOfLongestSubstring("abcdec")
	fmt.Println("length of longest substring:", length)

	median := algorithm.FindMedianSortedArrays([]int{1, 2}, []int{3, 5})
	fmt.Println("Median of Two Sorted Arrays:", median)
}
