package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AvineshTripathi/cred/analyze"
	"github.com/AvineshTripathi/cred/scan"
	"github.com/AvineshTripathi/cred/utils"
	"github.com/go-git/go-git/v5"
)

func main() {
	var input1 string // to store inputs  AKIA3XWNRUIJEF25GO7S
	// convert this to struct
	fmt.Println("Enter the repo Url: ")
	fmt.Scanln(&input1)


	nums := flag.Int("commits", 0, "Number of commits you want to check in")
	flag.Parse()

	// store results of the scan, convert this to struct
	var idRepo [][][]string
	var idCommit [][][]string
	var keyRepo [][][]string
	var keyCommit [][][]string

	f, err := os.MkdirTemp(os.TempDir(), "cred")
	if err != nil {
		fmt.Println(err)
	}

	defer os.RemoveAll(f)

	r, err := git.PlainClone(f, false, &git.CloneOptions{
		URL: input1,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want a latest repo scan(1) or commit scan(2)?")
	fmt.Scanln(&input1)

	// this will scan the entire commit history
	if input1 == "2" {
		commitIt, err := r.CommitObjects()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*nums)
		idCommit, keyCommit, err = scan.ReadCommits(commitIt, *nums)
		if err != nil {
			fmt.Println(err)
		}

		for _, i := range idCommit {
			for _, j := range keyCommit {
				t, err := analyze.Validate(i, j)
				if err != nil {
					fmt.Println(err)
				}

				utils.Printer(t)
			}
		}
	}

	// this will scan the latest code cloned temp in local
	if input1 == "1" {
 		idRepo, keyRepo, err = scan.ReadRepo(f)
		if err != nil {
			fmt.Println(err)
		}

		for _, i := range idRepo {
			for _, j := range keyRepo {
				t, err := analyze.Validate(i, j)
				if err != nil {
					fmt.Println(err)
				}

				utils.Printer(t)
			}
		}

	}
}
