package anybar

import (
	"fmt"
	"net"
)

// Commands to change dot color/status.
const (
	White       = "white"
	Red         = "red"
	Orange      = "orange"
	Yellow      = "yellow"
	Green       = "green"
	Cyan        = "cyan"
	Blue        = "blue"
	Purple      = "purple"
	Black       = "black"
	Question    = "question"
	Exclamation = "exclamation"
	Quit        = "quit"
)

// Command sends a command to an Anybar instance.
func Command(port int, cmd string) error {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write([]byte(cmd))
	_ = n
	return err
}
