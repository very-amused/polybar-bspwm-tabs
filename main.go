package main

import (
	"fmt"

	"github.com/very-amused/polybar-bspwm-tabs/bspc"
)

func main() {
	windows := bspc.QueryWindows()
	fmt.Println(FormatTabs(windows))
}
