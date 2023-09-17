package cmd

import (
	"fmt"
	"goping/lib"
	"goping/pkg"
	"math/rand"
	"net"
	"time"

	"github.com/urfave/cli/v2"
)

func PingCommand(ping *lib.CLICategory) *cli.Command {
	return &cli.Command{
		Name:    "ping",
		Usage:   "ping a host or cidr range",
		Aliases: []string{"p"},
		Action: func(c *cli.Context) error {
			arguments := c.Args().Slice()
			if err := pingAction(arguments); err != nil {
				return err
			}
			var input string
			fmt.Scanln(&input)
			return nil
		},
	}
}

// Takes a slice of strings and concurrently pings each valid host.
func pingAction(args []string) error {
	var c chan string = make(chan string)
	go printResult(c)

	// Ping all single hosts concurrently
	filterHostFunc := func(s string) bool { return net.ParseIP(s) != nil }
	for _, host := range pkg.FilterBy(args, filterHostFunc) {
		go pingHost(c, host)
	}

	filterCIDRFunc := func(s string) bool {
		_, _, err := net.ParseCIDR(s)
		if err != nil {
			return false
		}
		return true
	}

	// Ping all hosts in a CIDR concurrently
	for _, cidr := range pkg.FilterBy(args, filterCIDRFunc) {
		for _, host := range pkg.Hosts(cidr) {
			go pingHost(c, host.String())
		}
	}

	return nil
}

// Concurrently pings a single host
func pingHost(c chan string, host string) {
	reply := []string{"ok", "ok", "error"}
	latency := time.Duration(rand.Intn(250))
	time.Sleep(latency * time.Millisecond)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c <- fmt.Sprintf("Host %s replied with reason \"%s\" in %d ms", host, reply[r.Intn(len(reply))], latency)
}

func printResult(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}
