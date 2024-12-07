package solutions

func MaximumOddBinaryNumber(s string) string {
	zero, one := "", ""

	for _, v := range s {
		if v == '1' {
			one += "1"
		} else {
			zero += "0"
		}
	}

	if len(one) > 1 {
		return one[1:] + zero + one[:1]
	}

	return zero + one
}
