package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Init() {
	fmt.Println("MyCLI started. Type 'help' for commands. Type 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\r", "")
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "exit":
			fmt.Println("Bye!")
			return
		case "help":
			fmt.Println("Available commands: help, echo, exit")
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
		}
	}
}
