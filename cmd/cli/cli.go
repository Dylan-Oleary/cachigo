package cli

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Dylan-Oleary/cachigo/cmd/store"
	tcp "github.com/Dylan-Oleary/cachigo/tcp/client"
)

func Init(conn net.Conn) {
	store := store.GetCache()
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
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
			continue
		case "exit":
			fmt.Println("Bye!")
			return
		case "help":
			fmt.Println("Available commands: help, echo, exit")
			continue
		case "get":
			res, err := tcp.SendRequest(conn, &tcp.Request{Data: tcp.RequestData{Command: "get", Key: args[1]}})

			if err != nil {
				fmt.Println("Error", err)
				continue
			}

			fmt.Printf("%s\n", res.Message)
			continue
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

			res, err := tcp.SendRequest(conn, &tcp.Request{Data: tcp.RequestData{Command: "set", Key: args[1], Value: args[2]}})

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Outcome:", res.Success)
			continue
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			continue
		}
	}
}
