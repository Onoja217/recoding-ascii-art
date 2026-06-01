package generator

func GeneratePattern(c rune) []string {
	switch c {
	case 'A':
		return []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		}
	case 'Z':
		return []string{
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		}
	}

	if c >= 'A' && c <= 'Z' {
		return []string{
			"######",
			"#    #",
			"#    #",
			"#    #",
			"#    #",
			"#    #",
			"######",
			"      ",
		}
	}
	return []string{}
}