package builder

import "strings"

type ArtBuilder struct {
	texts []string
	style string
}

func NewArtBuilder() *ArtBuilder { return &ArtBuilder{style: "normal"} }

func (a *ArtBuilder) AddText(t string) *ArtBuilder {
	a.texts = append(a.texts, t)
	return a
}

func (a *ArtBuilder) SetStyle(s string) *ArtBuilder {
	if s != "normal" && s != "bold" && s != "italic" && s != "outline" {
		panic("invalid style")
	}
	a.style = s
	return a
}

func (a *ArtBuilder) Build() string {
	var out strings.Builder
	text := strings.Join(a.texts, "")

	for i := 0; i < 8; i++ {
		for _, c := range text {
			switch a.style {
			case "bold":
				out.WriteString(string(c));out.WriteString(string(c))
			case "italic":
				out.WriteString("/");out.WriteString(string(c))
			case "outline":
				out.WriteString("|");out.WriteString(string(c));out.WriteString("|")
			default:
				out.WriteRune(c)
			}
		}
		out.WriteString("\n")
	}
	return out.String()
}

