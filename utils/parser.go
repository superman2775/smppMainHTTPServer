package utils

import (
	"regexp"
	"strings"
)

func ParseMd(mdText string, page string) string {
	lines := strings.Split(mdText, "\n")
	var builder strings.Builder
	insideCard := false
	insideList := false
	firstCardMade := false
	firstCardOpen := false

	boldItalicRe := regexp.MustCompile(`\*\*\_(.*?)\_\*\*`) // **_text_**
	boldRe := regexp.MustCompile(`\*\*(.*?)\*\*`)           // **text**
	italicRe := regexp.MustCompile(`\_(.*?)\_`)             // _text_

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "##") {
			insideList = false
			// Close previous card if open
			if insideCard {
				builder.WriteString("</div>\n") // Close .card
				if firstCardOpen {
					builder.WriteString("</div>\n") // Close .first-card
					firstCardOpen = false
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
				if page == "roadmap" {
					builder.WriteString(`<div class="upcoming-update">Upcoming</div>` + "\n")
				}
				firstCardOpen = true
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
			insideList = false
			if insideCard {
				builder.WriteString("</div>\n")
				if firstCardOpen {
					builder.WriteString("</div>\n")
					firstCardOpen = false
				}
				insideCard = false
			}
			content := strings.TrimPrefix(line, "# ")
			builder.WriteString("<h1>" + content + "</h1>\n")

		} else if strings.HasPrefix(line, "IMG ") {
			builder.WriteString(`<div class=card-backdrop-img style="background-image: url(` + strings.TrimPrefix(line, "IMG ") + `)"> <div class=backdrop-card></div></div>`)
		} else if strings.HasPrefix(line, "ICO ") {
			builder.WriteString(`<div class=card-icon>` + strings.TrimPrefix(line, "ICO ") + `</div>`)
		} else if strings.HasPrefix(line, "- ") {
			line := strings.TrimPrefix(line, "- ")
			if !insideList {
				builder.WriteString("<ul>")
			}
			line = boldItalicRe.ReplaceAllString(line, "<strong><em>$1</em></strong>")
			line = boldRe.ReplaceAllString(line, "<strong>$1</strong>")
			line = italicRe.ReplaceAllString(line, "<em>$1</em>")
			builder.WriteString(`<li>` + line + `</li>`)
			insideList = true
		} else if line != "" {
			insideList = false
			line = boldItalicRe.ReplaceAllString(line, "<strong><em>$1</em></strong>")
			line = boldRe.ReplaceAllString(line, "<strong>$1</strong>")
			line = italicRe.ReplaceAllString(line, "<em>$1</em>")
			builder.WriteString("<p>" + line + "</p>\n")
		}
		if !insideList {
			builder.WriteString("</ul>")
		}
	}

	// Final card cleanup
	if insideCard {
		builder.WriteString("</div>\n")
		if firstCardOpen {
			builder.WriteString("</div>\n") // Close first-card wrapper
		}
	}
	if insideList {
		builder.WriteString("</ul>")
	}
	return builder.String()
}
