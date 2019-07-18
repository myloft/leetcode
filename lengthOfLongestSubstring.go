/**
 *Given a string, find the length of the longest substring without repeating characters.
 */

func lengthOfLongestSubstring(s string) int {
	var lens, head int
	for k, _ := range s {
		index := indexs(s[head:k], s[k])
		if index == -1 {
			if lens < k-head+1 {
				lens = k - head + 1
			}
		} else {
			head = head + index + 1
		}
	}
	return lens
}

func indexs(s string, b byte) int {
	for k, _ := range s {
		if s[k] == b {
			return k
		}
	}
	return -1
}