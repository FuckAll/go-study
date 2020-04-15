package lettcode

import (
	"testing"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

/*
	暴力法 O(n^3)
*/
func lengthOfLongestSubstringV1(s string) int {
	allUnique := func(runes []rune) bool {
		unique := make(map[rune]rune)
		for _, i := range runes {
			if _, ok := unique[i]; ok {
				return false
			} else {
				unique[i] = i
			}
		}
		return true
	}
	var byteArray = ([]rune) (s)
	max := 0
	for i := 0; i < len(byteArray); i++ {
		for j := i + 1; j <= len(byteArray); j++ {
			if allUnique(byteArray[i:j]) {
				if max < len(byteArray[i:j]) {
					max = len(byteArray[i:j])
				}
			}
		}
	}
	return max
}

func TestLengthOfLongestSubstringV1(t *testing.T) {
	v1 := lengthOfLongestSubstringV1("abcabcbb")
	if v1 != 3 {
		t.Error("v1 not equal 3")
	}

	v2 := lengthOfLongestSubstringV1("bbbbb")
	if v2 != 1 {
		t.Error("v2 not equal 1")
	}

	v3 := lengthOfLongestSubstringV1("pwwkew")
	if v3 != 3 {
		t.Error("v3 not equal 3")
	}
	v4 := lengthOfLongestSubstringV1("")
	if v4 != 0 {
		t.Error("v4 not equal 0")
	}
	v5 := lengthOfLongestSubstringV1(" ")
	if v5 != 1 {
		t.Errorf("v4 not equal 1 , get %d", v5)
	}

	v6 := lengthOfLongestSubstringV1("au")
	if v6 != 2 {
		t.Errorf("v6 not equal 2 , get %d", v6)
	}
}

/*
	滑动窗口
"" -> 0
" " -> 1
"abcabcabc"
*/
func lengthOfLongestSubstringV2(s string) int {
	max := 0
	runes := ([]rune) (s)
	j, i := 0, 0
	unique := make(map[rune]rune)
	for ; j < len(runes) && i < len(runes); {
		if _, ok := unique[runes[j]]; !ok {
			unique[runes[j]] = runes[j]
			j++
			if max < j-i {
				max = j - i
			}
		} else {
			delete(unique, runes[i])
			i++
		}
	}
	return max
}

func TestLengthOfLongestSubstringV2(t *testing.T) {
	v1 := lengthOfLongestSubstringV2("abcabcbb")
	if v1 != 3 {
		t.Errorf("v1 not equal 3,  get %d", v1)
	}

	v2 := lengthOfLongestSubstringV2("bbbbb")
	if v2 != 1 {
		t.Error("v2 not equal 1")
	}

	v3 := lengthOfLongestSubstringV2("pwwkew")
	if v3 != 3 {
		t.Error("v3 not equal 3")
	}
	v4 := lengthOfLongestSubstringV2("")
	if v4 != 0 {
		t.Error("v4 not equal 0")
	}
	v5 := lengthOfLongestSubstringV2(" ")
	if v5 != 1 {
		t.Errorf("v4 not equal 1 , get %d", v5)
	}

	v6 := lengthOfLongestSubstringV2("au")
	if v6 != 2 {
		t.Errorf("v6 not equal 2 , get %d", v6)
	}
}

/*
  优化滑动窗口(直接将i跳到重复的地方)
"" ->
" " ->
"abcabcabc"
"tmmzuxt"
*/
func lengthOfLongestSubstringV3(s string) int {
	max := 0
	runes := ([]rune) (s)
	unique := make(map[rune]int) // t 0  m 1 m
	for i, j := 0, 0; j < len(runes) && i < len(runes); j++ {
		if idx, ok := unique[runes[j]]; ok {
			if i < idx+1 {
				i = idx + 1
			}
		}
		unique[runes[j]] = j
		if max < j-i+1 {
			max = j - i + 1
		}
	}
	return max
}

func TestLengthOfLongestSubstringV3(t *testing.T) {
	v1 := lengthOfLongestSubstringV3("abcabcbb")
	if v1 != 3 {
		t.Errorf("v1 not equal 3,  get %d", v1)
	}

	v2 := lengthOfLongestSubstringV3("bbbbb")
	if v2 != 1 {
		t.Error("v2 not equal 1")
	}

	v3 := lengthOfLongestSubstringV3("pwwkew")
	if v3 != 3 {
		t.Error("v3 not equal 3")
	}
	v4 := lengthOfLongestSubstringV3("")
	if v4 != 0 {
		t.Error("v4 not equal 0")
	}
	v5 := lengthOfLongestSubstringV3(" ")
	if v5 != 1 {
		t.Errorf("v4 not equal 1 , get %d", v5)
	}

	v6 := lengthOfLongestSubstringV3("au")
	if v6 != 2 {
		t.Errorf("v6 not equal 2 , get %d", v6)
	}

	v7 := lengthOfLongestSubstringV3("tmmzuxt")
	if v7 != 5 {
		t.Errorf("v7 not equal 5, get %d", v7)
	}
}
