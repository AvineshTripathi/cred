package main

import (
	"fmt"
	"os"

	"github.com/AvineshTripathi/cred/scan"
	"github.com/go-git/go-git/v5"
)



func main() {
	var input1 string
	// convert this to struct
	fmt.Println("Enter the repo Url: ")
	fmt.Scanln(&input1)
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

	if input1 == "2" {
		commitIt, err := r.CommitObjects()
		if err != nil {
			fmt.Println(err)
		}

		idCommit, keyCommit, err = scan.ReadCommits(commitIt)
		if err != nil {
			fmt.Println(err)
		}

		for _, i := range idCommit {
			for _, j := range keyCommit {
				t, err := validate(i, j)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(t)
			}
		}
	}

	if input1 == "1" {

		idRepo, keyRepo, err = scan.ReadRepo(f)
		if err != nil {
			fmt.Println(err)
		}

		for _, i := range idRepo {
			for _, j := range keyRepo {
				t, err := validate(i, j)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(t)
			}
		}

	}
}
