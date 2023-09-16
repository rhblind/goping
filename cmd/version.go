package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:   "version",
		Usage:  "Show version information",
		Action: ShowVersionInfo,
	}
}

func ShowVersionInfo(c *cli.Context) error {
	fmt.Printf("Version: %s\n", c.App.Version)
	fmt.Printf("Build information: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	return nil
}
