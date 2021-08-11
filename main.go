package main

import (
	"cli/config"
	"flag"
)

func main() {
	c := config.NewConfig()
	c.Setup()

	flag.Parse()

	if err := c.ExecuteCommand(); err != nil {
		panic(err)
	}
}
