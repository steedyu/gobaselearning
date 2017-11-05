package littleexercise

import "math"

func ReverseIntNum(x int) (num int) {
	for x != 0 {
		t := x % 10
		x = x /10
		num = num * 10 + t
	}

	if math.MaxInt32 < num && math.MinInt32 > num {
		return 0
	}

	return
}
