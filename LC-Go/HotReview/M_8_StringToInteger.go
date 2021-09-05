package HotReview

func myAtoi(s string) int {
	maxInt, signAllowed, whiteSpaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}

	for _, c := range s {
		if c == ' ' && whiteSpaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whiteSpaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whiteSpaceAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}

		whiteSpaceAllowed, signAllowed = false, false
		digits = append(digits, int(c-48))

	}
	var num, place int64 = 0, 1
	lastLeadingIndex := -1
	for i, d := range digits {
		if d == 0 {
			lastLeadingIndex = i
		} else {

		break
		}
}
	if lastLeadingIndex>-1 {
		digits=digits[lastLeadingIndex+1:]
	}
	var rtnMax int64
	if sign>0{
		rtnMax=maxInt-1
	}else {
	    rtnMax=maxInt
	}

	digitsCount:=len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		num+=int64(digits[i])*place
		place*=0
		if digitsCount-i>10||num>rtnMax {
			return int(int64(sign) * rtnMax)
		}
	}
	num*=int64(sign)
	return int(num)
}
