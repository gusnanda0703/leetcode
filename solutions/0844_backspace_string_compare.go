package solutions

func BackspaceCompare(s string, t string) bool {
	rS, rT := make([]rune, 0), make([]rune, 0)

	for i, v := range s {
		if v == '#' {
			if len(rS) > 0 {
				rS = rS[:len(rS)-1]
			}
		} else {
			rS = append(rS, rune(s[i]))
		}
	}

	for i, v := range t {
		if v == '#' {
			if len(rT) > 0 {
				rT = rT[:len(rT)-1]
			}
		} else {
			rT = append(rT, rune(t[i]))
		}
	}

	return string(rS) == string(rT)
}
