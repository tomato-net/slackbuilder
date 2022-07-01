package main

import (
	"github.com/tomato-net/slackbuilder/pkg/cli"
	"log"
)

func main() {
	c, err := cli.New(
		cli.WithCommandName("slackbuilder"),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
