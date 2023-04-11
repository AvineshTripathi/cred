package scan

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AvineshTripathi/cred/analyze"
)

func ReadRepo(path string) ([][][]string, [][][]string, error) {

	var a [][][]string
	var b [][][]string

	err := filepath.WalkDir(path, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("h1")
			return err
		}
		if !info.IsDir() {

			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("h2")
				return err
			}

			matchedId, matchedkey := analyze.Find(string(content))

			if len(matchedId) != 0 {
				matchedId[0][0] = path
				fmt.Println(matchedId[0][0])
				a = append(a, matchedId)
			} else {
			}

			if len(matchedkey) != 0 {
				matchedkey[0][0] = path
				b = append(b, matchedkey)
			} else {

			}

		}

		return nil
	})
	if err != nil {
		return a, b, err
	}

	return a, b, nil
}