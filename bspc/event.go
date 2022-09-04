package bspc

// Event - An event from `bspc subscribe report`
type Event struct {
	Monitors []Monitor
}

// Monitor - A bspwm monitor
type Monitor struct {
	Name     string
	Active   bool
	Desktops []Desktop
}

// Desktop - A bspwm desktop
type Desktop struct {
	Name   string
	Active bool
}
