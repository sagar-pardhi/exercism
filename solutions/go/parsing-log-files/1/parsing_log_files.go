package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(?:\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	separator := regexp.MustCompile(`<[-~*=]*>`)

	// Split the line on the separator
	return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	quoteRegex := regexp.MustCompile(`"([^"]*)"`)

	count := 0

	for _, line := range lines {
		// Extract the quoted portion (if any)
		matches := quoteRegex.FindStringSubmatch(line)
		if len(matches) < 2 {
			continue // no quoted text found
		}

		quotedText := matches[1] // content inside quotes

		// Check if quoted text contains "password" (case-insensitive)
		if strings.Contains(strings.ToLower(quotedText), "password") {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)\s+`)

	result := make([]string, len(lines))

	for i, line := range lines {
		matches := re.FindStringSubmatch(line)

		if len(matches) < 2 {
			// No match → unchanged
			result[i] = line
			continue
		}

		username := matches[1]
		result[i] = fmt.Sprintf("[USR] %s %s", username, line)
	}

	return result
}
