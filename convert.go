//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var res string
	n := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			res += string(s[i+j])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				res += string(s[j+cycleLen-i])
			}
		}
	}
	return res
}