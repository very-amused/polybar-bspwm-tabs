package main

import (
	"fmt"
	"strings"

	"github.com/very-amused/polybar-bspwm-tabs/bspc"
)

type tabIndex struct {
	value      uint
	initialTab int // Index of the tab that resulted in the index being created, used to add a (0) indicator to tabs when they're found to be part of a set
}

func formatTabs() string {
	windows := bspc.QueryWindows()
	// If there are no monocle windows open, clear the module output
	if windows == nil || len(windows) < 2 {
		return ""
	}

	// Map the number of active windows under a given class to append numbers to duplicates
	indicators := make(map[string]*tabIndex)

	tabs := make([]string, len(windows)) // Inner tab cell content
	for i, w := range windows {
		tabs[i] = w.ClassName
		if indicators[w.ClassName] != nil {
			ind := indicators[w.ClassName]
			tabs[i] += fmt.Sprintf(" (%d)", ind.value)

			// If the previous tab was not known to be part of a set at the time, add its indicator
			if ind.value == 1 {
				tabs[ind.initialTab] += " (0)"
			}
		} else {
			indicators[w.ClassName] = &tabIndex{
				initialTab: i}
		}
		indicators[w.ClassName].value++
	}

	var out strings.Builder
	const border = '\u2571'
	out.WriteRune(border) // Left border
	for i, t := range tabs {
		if windows[i].Active {
			// Reverse fg/bg colors
			out.WriteString("%{R}")
		}
		out.WriteString(" " + t + " ")
		if windows[i].Active {
			out.WriteString("%{R}")
		}
		out.WriteRune(border)
	}
	return out.String()
}
