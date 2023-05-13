package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	out := strings.Builder{}

chars:
	for i := 0; i < len(strs[0]); i++ { // iterate over 1st word
		thisChar := strs[0][i]

		for j := 1; j < len(strs); j++ { // iterate over 2 to n words
			word := strs[j]
			if i == len(word) {
				break chars
			}

			if word[i] != thisChar {
				break chars
			}
		}

		out.WriteByte(thisChar)
	}

	return out.String()
}
