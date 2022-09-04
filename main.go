package main

import (
	"encoding/json"
	"fmt"

	"github.com/very-amused/polybar-bspwm-tabs/bspc"
)

func main() {
	windows := bspc.QueryWindows()
	m, _ := json.MarshalIndent(windows, "", "  ")
	fmt.Println(string(m))
}
