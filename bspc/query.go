package bspc

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
)

func QueryTabs() {
	// Check if bspc is a valid command
	bspc, err := exec.LookPath("bspc")
	if errors.Is(err, exec.ErrNotFound) {
		log.Fatal("bspc was not found in your PATH")
	} else if err != nil {
		panic(err)
	}

	cmd := exec.Command(bspc, "query",
		"-d", "focused.monocle",
		"-T")

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var bTree *Tree
	json.NewDecoder(stdout).Decode(&bTree)
	json.NewEncoder(os.Stdout).Encode(bTree)
}
