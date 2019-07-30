// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
func isPalindrome(x int) bool {
	reverse := 0
	temp := x
	if x < 0 {
		return false
	}
	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse == temp
}