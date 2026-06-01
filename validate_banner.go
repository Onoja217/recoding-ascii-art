package main

import (
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
	
	if len(banner) != 95 {
		return fmt.Errorf("banner has %d entries, expected 95", len(banner))
	}
	for char, slice := range banner {
		if char < 32 || char > 126 {
			return fmt.Errorf("invalid banner")
		}

		if len(slice) != 8 {
			return fmt.Errorf("invalid %c", char)
		}
	}
	return nil
}
