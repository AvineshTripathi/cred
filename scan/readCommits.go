package scan

import (
	"fmt"
	"log"

	"github.com/AvineshTripathi/cred/analyze"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// needs to be reconsider on the approch of iterating though all the commits
// we are using the go-git repo to get the iterator
// skip condition is still left to consider
func ReadCommits(commitIterator object.CommitIter, commits int) ([][][]string, [][][]string, error) {

	var a [][][]string
	var b [][][]string

	if commits == 0 {
		log.Println("No commits to verify")
		return nil, nil, nil
	} 
	err := commitIterator.ForEach(func(c *object.Commit) error {
		if commits <= 0 {
			return nil 
		}
		fileIt, err := c.Files()
		if err != nil {
			fmt.Println(err)
		}

		err = fileIt.ForEach(func(f *object.File) error {
			con, err := f.Contents()
			if err != nil {
				fmt.Println(err)
			}

			m, n := analyze.Find(con)

			if len(m) != 0 {

				a = append(a, m)
			}

			if len(n) != 0 {
				b = append(b, n)
			}

			return nil
		})
		commits--
		return nil
	})
	if err != nil {
		return a, b, err
	}

	return a, b, nil
}