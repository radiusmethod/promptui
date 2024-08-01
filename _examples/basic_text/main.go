package main

import (
	"fmt"
	"github.com/radiusmethod/promptui"
)

func main() {
	items := []string{"Option 1", "Option 2", "Option 3", "Option 4", "Option 5", "Option 6", "Option 7"}

	prompt := promptui.Select{
		Label:             "Select an Option",
		Items:             items,
		Size:              5,
		StartInSearchMode: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %q\n", result)
}
