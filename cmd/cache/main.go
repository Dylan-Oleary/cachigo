package main

import (
	"github.com/Dylan-Oleary/cachigo/cmd/cli"
)

func main() {
	c := cli.Init()
	c.Run()
}
