package solutions

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	rNum := 0
	for i := x; i > 0; i /= 10 {
		rNum = rNum*10 + i%10
	}

	return x == rNum
}
