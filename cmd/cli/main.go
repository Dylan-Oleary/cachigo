package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dylan-Oleary/cachigo/tcp"
)

func main() {
	conn, err := tcp.GetClient("localhost:8080")

	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

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
		case "keys":
			res, err := tcp.SendRequest(conn, &tcp.Request{Data: tcp.RequestData{Command: "keys"}})

			if err != nil {
				fmt.Println("Error", err)
				continue
			}

			fmt.Printf("%s\n", res.Message)
			continue
		case "del":
			res, err := tcp.SendRequest(conn, &tcp.Request{Data: tcp.RequestData{Command: "del", Key: args[1]}})

			if err != nil {
				fmt.Println("Error", err)
				continue
			}

			fmt.Printf("%s\n", res.Message)
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
