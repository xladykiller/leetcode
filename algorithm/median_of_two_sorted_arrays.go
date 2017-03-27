package algorithm

/**
https://leetcode.com/problems/median-of-two-sorted-arrays/
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
 */
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	mid := (length1 + length2) / 2
	store := make([]int, length1+length2)

	var i, j, k int = 0, 0, 0
	for i < length1 || j < length2 {
		if i < length1 && j < length2 {
			if nums1[i] > nums2[j] {
				store[k] = nums2[j]
				k++
				j++
			} else {
				store[k] = nums1[i]
				k++
				i++
			}
			continue
		}
		if j < length2 {
			store[k] = nums2[j]
			k++
			j++
			continue
		}
		if i < length1 {
			store[k] = nums1[i]
			k++
			i++
			continue
		}
	}

	if mid*2 == (length1 + length2) {
		return float64(store[mid]+store[mid-1]) / 2.0
	} else {
		return float64(store[mid])
	}

}
