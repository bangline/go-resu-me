package template_funcs

import (
	"fmt"
	"html/template"
	"regexp"
)

func InPairs(items []string) [][2]string {
	var result [][2]string
	for i := 0; i < len(items); i += 2 {
		pair := [2]string{items[i], ""}
		if i+1 < len(items) {
			pair[1] = items[i+1]
		}
		result = append(result, pair)
	}
	return result
}

func Markdown(text string) template.HTML {
	boldFormat := string(MarkupBoldText(text))
	return MarkupLinks(boldFormat)
}

func MarkupLinks(text string) template.HTML {
	re := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

	result := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}

		displayText := parts[1]
		url := parts[2]

		return fmt.Sprintf(`<a href="%s" class="underline">%s</a>`,
			template.HTMLEscapeString(url),
			template.HTMLEscapeString(displayText))
	})
	return template.HTML(result)
}

func MarkupBoldText(text string) template.HTML {
	re := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	result := re.ReplaceAllString(text, "<strong>$1</strong>")

	return template.HTML(result)
}
