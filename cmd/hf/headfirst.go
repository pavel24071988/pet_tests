package main

import (
	"strings"
)

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	if len(phrases) == 2 {
		result += " and "
	} else {
		result += ", and "
	}
	result += phrases[len(phrases)-1]
	return result
}
