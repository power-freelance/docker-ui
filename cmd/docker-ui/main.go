package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-ui"
	app.Version = "0.0.0"
	app.Description = "UI for docker"
	app.Author = "Memory Clutter <memclutter@gmail.com>"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
