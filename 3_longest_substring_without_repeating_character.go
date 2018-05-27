package main

import "fmt"

func max(num1, num2 int) int{
	if num1 > num2{
		return num1
	}else{
		return num2
	}
}
func lengthOfLongestSubstring(s string) int {
	/*
	i 记录字符串开始的位置
	j 记录末尾
	characterIndexMap存储字符的最大位置
	当出现重复的时候，直接把i挪到该字符的上次最大位置
	*/
	var characterIndexMap = map[byte]int{}
	ans := 0
	for i,j := 0, 0; j < len(s); j++{
		index, ok := characterIndexMap[s[j]]
		if ok && index >= i{
			i = index + 1	
		}else{
			ans = max(ans, j - i + 1)
		}
		characterIndexMap[s[j]] = j
	}
	return ans
}

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
