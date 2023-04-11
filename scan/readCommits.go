package scan

import (
	"fmt"

	"github.com/AvineshTripathi/cred/analyze"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// needs to be reconsider on the approch of iterating though all the commits
func ReadCommits(commitIterator object.CommitIter) ([][][]string, [][][]string, error) {

	var a [][][]string
	var b [][][]string
	err := commitIterator.ForEach(func(c *object.Commit) error {
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

		return nil
	})
	if err != nil {
		return a, b, err
	}

	return a, b, nil
}