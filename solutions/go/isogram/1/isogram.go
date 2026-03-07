package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	for i := range word {
		for j := i + 1; j < len(word); j++ {
			if word[i] == '-' || word[i] == ' ' {
				continue
			}
			if strings.EqualFold(string(word[i]), string(word[j])) {
				return false
			}
		}
	}
	return true
}
