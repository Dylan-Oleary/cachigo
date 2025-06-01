package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	tcp "github.com/Dylan-Oleary/cachigo/tcp/client"
)

var host = "localhost:8080"

// Next Steps
//// CLI -> TCP Client
//// TCP Client (Multiple) -> One TCP Server
//// Concurrency/MutEx/Profit?

func main() {
	ln, err := net.Listen("tcp", host)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer ln.Close()

	fmt.Printf("Server listening on %s\n", host)

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection")
			continue
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		buffer := make([]byte, 1024)
		size, err := c.Read(buffer)

		if err != nil {
			fmt.Println("Connection closed", err)
			break
		}

		data := tcp.GetRequest{}
		err = json.Unmarshal(buffer[:size], &data)

		if err != nil {
			fmt.Println("Invalid payload passed", err)
			continue
		}

		res := tcp.Response{Success: true, Message: "Your command was" + data.Command}
		b, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Failed to encode response", err)
			continue
		}

		c.Write(b)
	}
}
