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
	lines := SplitInput(input)
	var output []string
	for i, line := range lines {
		if line == "" {
			if i == len(lines)-1{
				output = append(output,RenderLine("",banner)...)
			}else {
			output = append(output, "")
			}
			continue
		}
		if _, err := ValidateInput(line); err != nil {
			return err.Error()
		}
		rendered := RenderLine(line, banner)
		output = append(output, rendered...)
	}
	return strings.Join(output, "\n") + "\n"
}
