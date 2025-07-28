package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/Dylan-Oleary/cachigo/store"
	"github.com/Dylan-Oleary/cachigo/tcp"
)

var host = "localhost:8080"

func main() {
	err := store.InitPersistence()
	if err != nil {
		os.Exit(1)
	}

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
		tcp.HandleRequest(&req, &res)

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
