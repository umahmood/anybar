package main

import (
	"flag"
	"log"

	"github.com/umahmood/anybar"
)

// Command line flags.
var cmd string
var port int

func init() {
	flag.StringVar(&cmd, "cmd", "", "The command to send to the Anybar instance.")
	// 1738 is the default port Anybar.app listens on.
	flag.IntVar(&port, "port", 1738, "The port Anybar is listening on.")
	flag.Parse()
}

func main() {
	err := anybar.Command(port, cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
