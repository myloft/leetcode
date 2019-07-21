// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := Max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]

}

func expandAroundCenter(s string, left int, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
