package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	tcp "github.com/Dylan-Oleary/cachigo/tcp/client"
	requests "github.com/Dylan-Oleary/cachigo/tcp/requests"
)

var host = "localhost:8080"

func main() {
	ln, err := net.Listen("tcp", host)

	if err != nil {
		fmt.Println("Failed to start server:", err)
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

		req := tcp.Request{}
		err = json.Unmarshal(buffer[:size], &req)

		if err != nil {
			fmt.Println("Invalid payload passed", err)
			continue
		}

		res := tcp.Response{}
		requests.HandleRequest(&req, &res)

		b, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Failed to encode response", err)
			res.Message = err.Error()
			res.Success = false
			break
		}

		c.Write(b)
	}
}
