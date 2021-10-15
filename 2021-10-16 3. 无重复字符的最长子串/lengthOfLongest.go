package main

import "fmt"

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Print(res)
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	l := len(s)
	rk, ans := -1, 0

	for i := 0; i < l; i++ {

		if i > 0 {
			delete(m, s[i-1])
		}

		for rk+1 < l && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		ans = max(ans, rk-i+1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
