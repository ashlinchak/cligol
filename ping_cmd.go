package main

import "fmt"

// PingCmd is the ping command handler
type PingCmd struct{}

// Pong resonse on helth check
func (ping PingCmd) Pong() {
	fmt.Println("PONG")
}
