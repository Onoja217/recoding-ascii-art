package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage : go run . <text>")
	}
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	text := os.Args[1]
	result := GenerateArt(text, banner)
	fmt.Print(result)
}
