package utils

import (
	"regexp"
	"strings"
)

func ParseMd(mdText string) string {
	lines := strings.Split(mdText, "\n")
	var builder strings.Builder
	var insideCard bool

	boldRe := regexp.MustCompile(`\*\*(.*?)\*\*`)
	italicRe := regexp.MustCompile(`\*(.*?)\*`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "## ") {
			if insideCard {
				insideCard = false
				builder.WriteString("</div>\n")
			}
			content := strings.TrimPrefix(line, "## ")
			builder.WriteString("<div class=card>\n")

			builder.WriteString("<h2>" + content + "</h2>\n")
			insideCard = true
		} else if strings.HasPrefix(line, "# ") {
			content := strings.TrimPrefix(line, "# ")
			builder.WriteString("<h1>" + content + "</h1>\n")
			if insideCard {
				insideCard = false
				builder.WriteString("</div>\n")
			}
		} else if strings.HasPrefix(line, "IMG ") {
			builder.WriteString(`<img class=backdropCard src="` + strings.TrimPrefix(line, "IMG ") + `">`)
		} else if line != "" {

			line = boldRe.ReplaceAllString(line, "<strong>$1</strong>")
			line = italicRe.ReplaceAllString(line, "<em>$1</em>")
			builder.WriteString("<p>" + line + "</p>\n")
		}
	}

	return builder.String()
}
