package main

func find(data string) ([][]string, [][]string) {
	matchedId := idPat.FindAllStringSubmatch(data, -1)
	matchedKey := secretPat.FindAllStringSubmatch(data, -1)

	return matchedId, matchedKey
}