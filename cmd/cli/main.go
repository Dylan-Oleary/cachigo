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
		case "del", "get":
			if len(args) != 2 {
				fmt.Println("Error: Invalid number of arguments passed")
				continue
			}

			data := tcp.RequestData{Command: args[0], Key: args[1]}
			req := tcp.Request{Data: data}
			res, err := tcp.SendRequest(conn, &req)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("%s\n", res.Message)
			continue
		case "keys":
			if len(args) != 1 {
				fmt.Println("Error: Invalid number of arguments passed")
				continue
			}

			data := tcp.RequestData{Command: args[0]}
			req := tcp.Request{Data: data}
			res, err := tcp.SendRequest(conn, &req)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("%s\n", res.Message)
			continue
		case "set":
			if len(args) != 3 {
				fmt.Println("Error: Invalid number of arguments passed")
				continue
			}

			data := tcp.RequestData{Command: args[0], Key: args[1], Value: args[2]}
			req := tcp.Request{Data: data}
			res, err := tcp.SendRequest(conn, &req)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("%s\n", res.Message)
			continue
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			continue
		}
	}
}
