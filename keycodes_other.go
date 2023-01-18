//go:build !windows
// +build !windows

package promptui

import "github.com/chzyer/readline"

var (
	// KeyBackspace is the default key for deleting input text.
	KeyBackspace rune = readline.CharBackspace

	// KeySpace is the default key to chose options for checkbox
	KeySpace rune = 32
)
