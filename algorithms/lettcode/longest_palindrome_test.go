package lettcode

import "testing"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
/*
	暴力枚举
*/
func longestPalindromeV1(s string) string {
	isPalindrome := func(subRunes []rune) bool {
		for i := 0; i < len(subRunes)/2; i++ {
			if subRunes[i] != subRunes[len(subRunes)-i-1 ] {
				return false
			}
		}
		return true
	}

	runes := ([]rune) (s)
	var ret []rune
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j <= len(runes); j++ {
			if isPalindrome(runes[i:j]) && len(ret) < len(runes[i:j]) {
				ret = runes[i:j]
			}
		}
	}
	return (string)(ret)
}

func TestLongestPalindrome(t *testing.T) {
	v1 := longestPalindromeV1("cbbd")
	if v1 != "bb" {
		t.Errorf("v1 want bb, but get %s\n", v1)
	}

	v2 := longestPalindromeV1("babad")
	if v2 != "bab" {
		t.Errorf("v2 want aba, but get %s\n", v2)
	}

}
