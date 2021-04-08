package mrz

import (
	"regexp"
	"strconv"
)

func clearText(str string) string {
	r := regexp.MustCompile("<+$")
	str = r.ReplaceAllString(str, "")
	r2 := regexp.MustCompile("<")
	return r2.ReplaceAllString(str, " ")
}

func getText(lines []string, pos *Position) string {
	return lines[pos.Line][pos.Start:pos.End]
}

func checkDigitCode(src, code string) bool {
	c := 0
	factors := []int{7, 3, 1}
	for i, charCode := range []rune(src) {
		code := int(charCode)
		if code == 60 {
			code = 0
		}
		if code >= 65 {
			code -= 55
		}
		if code >= 48 {
			code -= 48
		}
		code *= factors[i%3]
		c += code
	}
	return strconv.Itoa(c%10) == code
}
