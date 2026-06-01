package main

import "strings"

func PadArtRows(rows []string, width int) []string {

	if width <= 0 {
		return rows
	}
	res := make([]string, len(rows))
	for i, ch := range rows {
		padding := width - len(ch)
		if padding > 1 {
			res[i] = ch + strings.Repeat(" ", padding)
		} else {
			res[i] = ch
		}
	}
	return res
}
