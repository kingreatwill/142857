package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring(" "))
	fmt.Printf("%v\n", lengthOfLongestSubstring("b"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("bb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("nfpdmpi"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 判断当前字符在积累的窗口位置.
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			// 没有找到:end = i+1
			if i+1 > end {// 这样就不用记录最大长度了。
				end = i + 1
			}
		} else {
			// 找到了，说明有重复:start和end 同步后移
			start = start + index + 1
			end = end + index + 1
		}
	}
	return end - start
}
