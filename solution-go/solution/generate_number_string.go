package solution

import (
	"errors"
	"fmt"
	"strings"
)

/*
Input = “dani Maulana”
Output = “d4a2nimul”
-> Format [A-z][occurrence of 2nd char][A-z][occurrence of third char][[a-Z]...:strip(" ", 2nd char,
3rd char)]]
*/
func GenerateNumberString(s string) (string, error) {
	s = strings.ToLower(s)

	if len(s) < 3 {
		return "", errors.New("string must be longer than 3 characters")
	}

	var sb strings.Builder
	ch1Occ := 1
	ch2Occ := 1
	sRune := []rune(s)
	for _, ch := range s[3:] {
		if ch == sRune[1] {
			ch1Occ += 1
			continue
		} else if ch == sRune[2] {
			ch2Occ += 1
			continue
		} else if ch == rune(' ') {
			continue
		}

		sb.WriteString(strings.ToLower(string(ch)))
	}

	return fmt.Sprintf("%s%d%s%d%s",
		string(s[0]),
		ch1Occ, string(s[1]),
		ch2Occ, string(s[2])) + sb.String(), nil
}
