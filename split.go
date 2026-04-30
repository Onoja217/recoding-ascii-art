package main

import (
	"strings"
)

func SplitLine(input string) []string {
	return strings.Split(input, "\\n")
}
