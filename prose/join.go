package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	phraseLength := len(phrases)
	if phraseLength == 0 {
		return ""
	}
	if phraseLength == 1 {
		return phrases[0]
	}
	if phraseLength == 2 {
		return phrases[0] + " and " + phrases[1]
	}
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
