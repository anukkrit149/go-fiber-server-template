package helpers

import (
	"encoding/json"
	"strings"
)

// StudlyCase converts foo-bar to FooBar.
// Reference: https://github.com/stoewer/go-strcase/blob/master/camel.go
func StudlyCase(s string) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	var prev rune
	for _, curr := range s {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || prev == 0 {
				buffer = append(buffer, toUpper(curr))
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
		prev = curr
	}

	return string(buffer)
}

func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || ch == ' '
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func ConvertMapToString(data map[string]string) string {
	empData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(empData)
}
