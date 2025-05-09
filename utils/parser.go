package utils

import (
	"regexp"
	"strings"
)

func ParseMd(mdText string) string {
	lines := strings.Split(mdText, "\n")
	var builder strings.Builder
	insideCard := false
	firstCardMade := false

	boldRe := regexp.MustCompile(`\*\*(.*?)\*\*`)
	italicRe := regexp.MustCompile(`\*(.*?)\*`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "##") {
			// Close previous card if open
			if insideCard {
				builder.WriteString("</div>\n")
				if firstCardMade {
					// Close wrapper for first card
					builder.WriteString("</div>\n")
				}
				insideCard = false
			}

			// Extract and clean title
			content := strings.TrimPrefix(line, "##")
			isLight := false
			isDark := false
			if strings.HasPrefix(content, " light") {
				isLight = true
				content = strings.TrimPrefix(content, " light")
			} else if strings.HasPrefix(content, " dark") {
				isDark = true
				content = strings.TrimPrefix(content, " dark")
			}

			// If this is the first card, open wrapper and add label
			if !firstCardMade {
				builder.WriteString(`<div class="first-card">` + "\n")
				builder.WriteString(`<div class="upcoming-update">Upcoming</div>` + "\n")
				firstCardMade = true
			}

			// Start the card itself
			if isLight {
				builder.WriteString(`<div class="card light">` + "\n")
			} else if isDark {
				builder.WriteString(`<div class="card dark">` + "\n")
			} else {
				builder.WriteString(`<div class="card">` + "\n")
			}

			builder.WriteString("<h2>" + content + "</h2>\n")
			insideCard = true

		} else if strings.HasPrefix(line, "# ") {
			if insideCard {
				builder.WriteString("</div>\n")
				if firstCardMade {
					builder.WriteString("</div>\n")
				}
				insideCard = false
			}
			content := strings.TrimPrefix(line, "# ")
			builder.WriteString("<h1>" + content + "</h1>\n")

		} else if strings.HasPrefix(line, "IMG ") {
			builder.WriteString(`<div class=backdrop-card-img style="background-image: url(` + strings.TrimPrefix(line, "IMG ") + `)"> <div class=backdrop-card></div></div>`)

		} else if line != "" {
			line = boldRe.ReplaceAllString(line, "<strong>$1</strong>")
			line = italicRe.ReplaceAllString(line, "<em>$1</em>")
			builder.WriteString("<p>" + line + "</p>\n")
		}
	}

	// Final card cleanup
	if insideCard {
		builder.WriteString("</div>\n")
		if firstCardMade {
			builder.WriteString("</div>\n") // Close first-card wrapper
		}
	}

	return builder.String()
}
