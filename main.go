package main

import (
	"fmt"
	"os"

	"goping/cmd"
	"goping/lib"

	"github.com/urfave/cli/v2"
)

const cliVersion = "0.0.1"
const (
	cliName         = "goping"
	cliUsage        = "A concurrent ping machine"
	cliAuthorName   = "Rolf Håvard Blindheim"
	cliAuthorEmail  = "rhblind@gmail.com"
	cliDescription1 = "goping is a ping machine that can ping multiple hosts concurrently.\n"
	cliDescription2 = "It is an example project for me to learn how to do concurrency in Go."
)

var (
	pingCategory = lib.CLICategory{
		Name: "ping",
	}
)

func main() {
	app := &cli.App{
		Name:        cliName,
		Version:     cliVersion,
		Usage:       cliUsage,
		Description: cliDescription1 + cliDescription2,
		Authors: []*cli.Author{
			{
				Name:  cliAuthorName,
				Email: cliAuthorEmail,
			},
		},
		Commands: []*cli.Command{
			cmd.VersionCommand(),
			cmd.PingCommand(&pingCategory),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("\033[31m%s\033[0m\n", err.Error())
	}
}
