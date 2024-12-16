package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestSubStrMatch(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
		{"a", false},
		{"aaaa", true},
		{"abcd", false},
	}

	for _, test := range tests {
		result := repeatedSubStrPattern(test.s)
		fmt.Printf("Input: %s\nOutput: %v\nExpected: %v\n\n", test.s, result, test.expected)
	}
}

func repeatedSubStrPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	doubled := s + s                      // 拼接两次
	subStr := doubled[1 : len(doubled)-1] // 去掉首尾字符

	// 如果原字符串是通过子串重复构成的，它一定是子串的一部分
	return strings.Contains(subStr, s)
}
