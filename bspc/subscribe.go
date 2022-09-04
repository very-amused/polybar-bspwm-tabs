package bspc

import (
	"bufio"
	"os/exec"
)

// Subscribe - Return a channel that blocks until bspwm node state has been updated, and should be queried again
func Subscribe() (c chan bool) {
	bspc := getBSPC()
	cmd := exec.Command(bspc, "subscribe",
		"desktop_layout", "node_focus") // Events to be notified on
	c = make(chan bool)
	stdout, _ := cmd.StdoutPipe()
	lines := bufio.NewScanner(stdout)
	cmd.Start()
	go func() {
		for lines.Scan() {
			c <- true
		}
	}()
	return c
}
