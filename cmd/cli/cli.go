package cli

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func Init() *prompt.Prompt {
	p := prompt.New(executor, completer)

	return p
}

func executor(in string) {
	in = strings.TrimSpace(in)
	args := strings.Split(in, " ")

	switch args[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "help":
		fmt.Println("Commands: help, echo, exit")
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "help", Description: "Show help"},
		{Text: "echo", Description: "Echo input"},
		{Text: "exit", Description: "Exit the CLI"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
