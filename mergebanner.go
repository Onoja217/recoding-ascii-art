package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)

	for ok, ch := range base {
		result[ok] = ch
	}
	for ok, ch := range priority {
		result[ok] = ch
	}
	return result
}
