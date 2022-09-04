package bspc

import (
	"encoding/json"
	"errors"
	"log"
	"os/exec"
)

// QueryWindows - Return a list of windows if the layout is monocle
func QueryWindows() (windows []Window) {
	// Check if bspc is a valid command
	bspc, err := exec.LookPath("bspc")
	if errors.Is(err, exec.ErrNotFound) {
		log.Fatal("bspc was not found in your PATH")
	} else if err != nil {
		panic(err)
	}

	// Decode a json tree from bspc
	cmd := exec.Command(bspc, "query",
		"-d", "focused.monocle",
		"-T")
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var tree *Tree
	json.NewDecoder(stdout).Decode(&tree)
	if tree == nil {
		return nil
	}

	windows = tree.Windows()
	// Set the active window from the tree's focused node ID
	for i, w := range windows {
		if w.nodeID == tree.FocusedNodeID {
			windows[i].Active = true
			break // There can only be 1 active window
		}
	}
	return windows
}
