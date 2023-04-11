package utils

import "fmt"

func Printer(t *[]Truth) {
	for _, ele := range *t {
		fmt.Println(ele.Found, ele.Id, ele.IdPath, ele.Key, ele.KeyPath)
	}
}