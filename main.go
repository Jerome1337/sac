package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "sac",
		Usage:                "Geometric shapes helper command!",
		Action:               prompt,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
