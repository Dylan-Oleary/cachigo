package main

import (
	"os"

	"github.com/Dylan-Oleary/cachigo/cmd/cli"
	tcp "github.com/Dylan-Oleary/cachigo/tcp/client"
)

func main() {
	conn, err := tcp.GetClient("localhost:8080")

	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	cli.Init(conn)
}
