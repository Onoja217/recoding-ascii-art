package main

import "strings"

func TrimArtRows(rows []string) []string {
	res := make([]string, len(rows))
	
	for i, ch := range rows {
		if strings.TrimSpace(ch) == "" {
			res[i] = ""
		}
		res[i] = strings.TrimRight(ch, " ")
	}
	return res
}
