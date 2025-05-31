package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dylan-Oleary/cachigo/cmd/store"
)

func Init() {
	store := store.InitCache()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Failed to read input.")
			continue
		}

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
		case "get":
			v, err := store.Get(args[1])

			if err != nil {
				fmt.Println("Error", err)
				continue
			}

			fmt.Printf("%s\n", v)
		case "list":
			fmt.Print("\n")

			for _, k := range store.ListKeys() {
				fmt.Println(k)
			}
			continue
		case "remove":
			store.Remove(args[1])
			continue
		case "set":
			if len(args) != 3 {
				fmt.Println("Invalid args provided")
				continue
			}

			_, err := store.Set(args[1], args[2])

			if err != nil {
				fmt.Println("Error", err)
			}
			continue
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
		}
	}
}
