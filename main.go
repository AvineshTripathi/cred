package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/go-git/go-git/v5"
)

var (
	idPat     = regexp.MustCompile(`\b((?:AKIA|ABIA|ACCA|ASIA)[0-9A-Z]{16})\b`)
	secretPat = regexp.MustCompile(`[^A-Za-z0-9+\/]{0,1}([A-Za-z0-9+\/]{40})[^A-Za-z0-9+\/]{0,1}`)
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

		idCommit, keyCommit, err = ReadCommits(commitIt)
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

		idRepo, keyRepo, err = ReadRepo(f)
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
