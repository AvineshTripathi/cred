package analyze

import "regexp"

var (
	idPat     = regexp.MustCompile(`\b((?:AKIA|ABIA|ACCA|ASIA)[0-9A-Z]{16})\b`)
	secretPat = regexp.MustCompile(`[^A-Za-z0-9+\/]{0,1}([A-Za-z0-9+\/]{40})[^A-Za-z0-9+\/]{0,1}`)
)

func Find(data string) ([][]string, [][]string) {
	matchedId := idPat.FindAllStringSubmatch(data, -1)
	matchedKey := secretPat.FindAllStringSubmatch(data, -1)

	return matchedId, matchedKey
}