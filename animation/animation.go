package animation

import (
	"fmt"
	"strings"
)

type Animation struct {
	text   string
	frames []string
}

func NewAnimation(t string, n int) *Animation {
	return &Animation{t, make([]string, n)}
}

func (a *Animation) GetFrame(i int) string {
	if len(a.frames) == 0 {
		return ""
	}
	return a.frames[i%len(a.frames)]
}

func build(lines []string) string {
	var b strings.Builder
	for i := 0; i < 10; i++ {
		if i < len(lines) {
			b.WriteString(lines[i])
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (a *Animation) GenerateSpinFrames() {
	r := []rune(a.text)

	for i := range a.frames {
		l := ""
		for j, c := range r {
			switch (i + j) % 4 {
			case 0:
				l += string(c)
			case 1:
				l += string(c) + string(c)
			case 2:
				l += "/" + string(c)
			default:
				l += "|" + string(c)
			}
		}

		lines := make([]string, 10)
		for k := range lines {
			lines[k] = l
		}
		a.frames[i] = build(lines)
	}
}

func (a *Animation) GenerateWaveFrames() {
	r := []rune(a.text)

	for i := range a.frames {
		lines := make([]string, 10)
		for j := 0; j < 10; j++ {
			lines[j] = strings.Repeat(" ", (i+j)%3) + string(r)
		}
		a.frames[i] = build(lines)
	}
}

func (a *Animation) GenerateZoomFrames() {
	r := []rune(a.text)

	for i := range a.frames {
		s := (i % 3) + 1
		lines := make([]string, 10)
		for j := 0; j < 10; j++ {
			lines[j] = strings.Repeat(string(r)+" ", s)
		}
		a.frames[i] = build(lines)
	}
}

func (a *Animation) Play() string {
	var b strings.Builder
	for i := range a.frames {
		b.WriteString(fmt.Sprintf("Frame %d\n%s", i, a.frames[i]))
	}
	return b.String()
}
