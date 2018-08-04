package medium

import "math"

func MyPow(x float64, n int) float64 {

	if x > math.MaxFloat64 || x < - math.MaxFloat64 {
		return  x
	}

	if x == 1 || x == -1 {
		if n % 2 == 0 {
			return 1
		} else {
			return x
		}
	}

	if n == 0 {
		return 1
	} else {
		if n < 0 {
			x = 1 / x
			n = -1 * n
		}

		ret := MyPow(x, n / 2)
		ret *= ret
		if n % 2 > 0 {
			ret *= x
		}

		return ret
	}
}