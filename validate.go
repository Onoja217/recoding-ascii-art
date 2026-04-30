package main

import "fmt"

func ValidateAll(s string) (rune, error) {
	if len(s) == 0 {
		return 0, nil
	}
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			return ch, fmt.Errorf("unsupported input%q", ch)
		}
	}
	return 0, nil
}
