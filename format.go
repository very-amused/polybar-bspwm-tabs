package main

import (
	"fmt"
	"os"
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
		tabs[i] = w.InstanceName
		if indicators[w.InstanceName] != nil {
			ind := indicators[w.InstanceName]
			tabs[i] += fmt.Sprintf(" (%d)", ind.value)

			// If the previous tab was not known to be part of a set at the time, add its indicator
			if ind.value == 1 {
				tabs[ind.initialTab] += " (0)"
			}
		} else {
			indicators[w.InstanceName] = &tabIndex{
				initialTab: i}
		}
		indicators[w.InstanceName].value++
	}

	var out strings.Builder
	const border = 0x2f
	for i, t := range tabs {
		out.WriteRune(' ')
		if windows[i].Active {
			bold := os.Getenv("BOLD_FONT")
			out.WriteString(fmt.Sprintf("%%{T%s}", bold))
		}
		out.WriteString(t)
		if windows[i].Active {
			normal := os.Getenv("FONT")
			out.WriteString(fmt.Sprintf("%%{T%s}", normal))
		}
		out.WriteRune(' ')
		if i < len(tabs)-1 {
			out.WriteRune(border)
		}
	}
	return out.String()
}
