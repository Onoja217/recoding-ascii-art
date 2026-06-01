package converter

import "testing"

func TestGenerateFont(t *testing.T) {
	font := GenerateFont()

	// must contain all printable ASCII chars
	if len(font) != 95 {
		t.Fatalf("expected 95 chars, got %d", len(font))
	}

	for c := rune(32); c < 127; c++ {
		art, ok := font[c]

		// character exists
		if !ok {
			t.Errorf("missing char %q", c)
			continue
		}

		// must have 8 lines
		if len(art) != 8 {
			t.Errorf("%q should have 8 lines, got %d", c, len(art))
		}

		for i, line := range art {
			// each line must be width 8
			if len(line) != 8 {
				t.Errorf("%q line %d width = %d, want 8", c, i, len(line))
			}

			// only valid chars allowed
			for _, r := range line {
				if r != '*' && r != '.' && r != ' ' {
					t.Errorf("%q contains invalid char %q", c, r)
				}
			}
		}
	}
}

func TestGenerateFont_UniqueChars(t *testing.T) {
	font := GenerateFont()

	a := ""
	b := ""

	for _, s := range font['A'] {
		a += s
	}
	for _, s := range font['B'] {
		b += s
	}

	if a == b {
		t.Error("A and B should not look identical")
	}
}

func TestGenerateFont_Space(t *testing.T) {
	font := GenerateFont()

	art := font[' ']

	if len(art) != 8 {
		t.Fatal("space should have 8 lines")
	}
}
