package main

import (
	"fmt"

	"github.com/very-amused/polybar-bspwm-tabs/bspc"
)

func main() {
	fmt.Println(formatTabs())
	update := bspc.Subscribe()
	for <-update {
		fmt.Println(formatTabs())
	}
}
