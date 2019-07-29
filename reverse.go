// Given a 32-bit signed integer, reverse digits of an integer.
const INT_MAX = int(^uint32(0) >> 1)
const Int_MIN = -INT_MAX - 1

func reverse(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if res > INT_MAX/10 || res == INT_MAX/10 && pop > 7 {
			return 0
		}
		if res < Int_MIN/10 || res == Int_MIN/10 && pop < -8 {
			return 0
		}
		res = res*10 + pop
	}
	return res
}
