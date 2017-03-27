package algorithm

/**
https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	result := make([]string, len(s))
	for i, c := range s {

		before := ""

		if i != 0 {
			before = result[i-1]
		}

		if before == "" {
			result[i] = string(c)
			continue
		}
		nstr := string(c)
		for j := len(before) - 1; j >= 0; j-- {

			if before[j] != byte(c) {
				nstr = string(before[j]) + nstr
			} else {
				result[i] = nstr
				break;
			}
		}
		result[i] = nstr

	}
	max := 1
	for _, s := range result {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}
