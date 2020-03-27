package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func formatCommas(num int) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1.$2")
	}
	return str
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
