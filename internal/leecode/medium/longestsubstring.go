package medium

import (
	"strings"
	"fmt"
)

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

func lengthOfLongestSubstring(s string) int {

	length := len(s)

	if length <= 1 {
		return len(s)
	}

	var maxlength int = -1
	var startindex int = 0
	var endindex int = 1
	for endindex < length {

		c := s[endindex]
		existsIndex := strings.IndexByte(s[startindex:endindex], c)
		if existsIndex == -1 {
			endindex ++
		}else {
			if endindex - startindex > maxlength {
				maxlength = endindex - startindex
			}
			startindex = existsIndex + startindex + 1
			endindex ++
		}
	}

	if endindex - startindex > maxlength {
		maxlength = endindex - startindex
	}

	return maxlength
}

func LengthSubstring() {

	//fmt.Println(lengthOfLongestSubstring("dvdf"))
	//fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
