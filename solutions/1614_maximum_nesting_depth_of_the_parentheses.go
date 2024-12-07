package solutions

func MaxDepth(s string) int {

	maxVps, vps := 0, 0

	for _, v := range s {
		if v == '(' {
			vps++
		} else if v == ')' {
			vps--
		}

		if vps < 0 {
			break
		}

		maxVps = max(maxVps, vps)
	}

	return maxVps
}
