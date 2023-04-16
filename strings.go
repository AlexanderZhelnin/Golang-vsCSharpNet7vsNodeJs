package main

func Compare(s1, s2 string) int {
	len1, len2 := len(s1), len(s2)
	i1, i2 := 0, 0

	for i1 < len1 && i2 < len2 {
		c1, c2 := s1[i1], s2[i2]

		if isDigit(c1) && isDigit(c2) {
			num1, newI1 := parseNumber(s1, i1)
			num2, newI2 := parseNumber(s2, i2)

			if num1 != num2 {
				if num1 < num2 {
					return -1
				}
				return 1
			}
			i1, i2 = newI1, newI2
		} else {
			if c1 != c2 {
				if c1 < c2 {
					return -1
				}
				return 1
			}
			i1, i2 = i1+1, i2+1
		}
	}

	if i1 < len1 {
		return 1
	} else if i2 < len2 {
		return -1
	} else {
		return 0
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseNumber(s string, i int) (int, int) {
	num := 0
	for i < len(s) && isDigit(s[i]) {
		num = 10*num + int(s[i]-'0')
		i++
	}
	return num, i
}
