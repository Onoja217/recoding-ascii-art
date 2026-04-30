package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {

	if input == "" {
		return ""
	}
	if strings.ReplaceAll(input, "\\n", "") == "" {
		return strings.Repeat("\n", len(input)/2)
	}
	lines := SplitLine(input)
	var output []string
	for _, line := range lines {
		if line == "" {
			output = append(output, "")
			continue
		}
		if _, err := ValidateAll(line); err != nil {
			return err.Error()
		}
		rendered := RenderLine(line, banner)
		output = append(output, rendered...)
	}
	return strings.Join(output, "\n") + "\n"
}
