package utils

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func Printer(t *[]Truth) {

	for _, ele := range *t {

		item := pterm.LeveledList{
			pterm.LeveledListItem{Level: 0, Text: fmt.Sprintf("Match - %t", ele.Found)},
			pterm.LeveledListItem{Level: 1, Text: "id"},
			pterm.LeveledListItem{Level: 2, Text: fmt.Sprintf("id - %s", ele.Id)},
			pterm.LeveledListItem{Level: 2, Text: fmt.Sprintf("Path - %s", ele.IdPath)},
			pterm.LeveledListItem{Level: 1, Text: "key"},
			pterm.LeveledListItem{Level: 2, Text: fmt.Sprintf("key - %s", ele.Key)},
			pterm.LeveledListItem{Level: 2, Text: fmt.Sprintf("Path - %s", ele.KeyPath)},
		}

		root := putils.TreeFromLeveledList(item)

		pterm.DefaultTree.WithRoot(root).Render()
	}

}