package solution

func ReverseString(s string) string {
	l := 0
	r := len(s) - 1

	sByte := []byte(s)

	for l < r {
		sByte[l], sByte[r] = sByte[r], sByte[l]
		l += 1
		r -= 1
	}

	return string(sByte)
}
