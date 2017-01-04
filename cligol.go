package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ping":
		pingCmd := PingCmd{}
		pingCmd.Pong()
	case "server":
		serverCmd := ServerCmd{}
		serverCmd.Init()
		serverCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
