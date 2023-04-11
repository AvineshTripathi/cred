package scan

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AvineshTripathi/cred/analyze"
)

// this func does the walk in the temp dir created where the cloned repo is stored
// initial approach was to read similar to commits appraoch however iterating the entire repo is still crucial thing
func ReadRepo(path string) ([][][]string, [][][]string, error) {

	var a [][][]string
	var b [][][]string

	err := filepath.WalkDir(path, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("h1")
			return err
		}

		// the name condition is specific to a particular demo repo, however there should a way to find those file which are satic and not need to be scanned like this
		// we can scan this however this would take a lot time doing
		// one workaround found over here was to use goroutine which is better but still needs a lot improvement inorder to remove this name check
		if !info.IsDir() && info.Name() != "package-lock.json" {

			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			matchedId, matchedkey := analyze.Find(string(content))

			if len(matchedId) != 0 {
				matchedId[0][0] = path
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