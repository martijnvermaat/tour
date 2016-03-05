package tour

import "strings"

func WordCount(s string) map[string]int {
	counts := make(map[string]int)

	for _, word := range strings.Fields(s) {
		counts[word] += 1
	}

	return counts
}
