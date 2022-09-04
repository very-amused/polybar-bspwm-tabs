package bspc

// A binary tree of nodes
type Tree struct {
	FocusedNodeID uint  `json:"focusedNodeId"` // The ID of the focused node
	Root          *Node `json:"root"`
}

// Windows - Return a flattened list of windows contained in a desktop tree
func (t Tree) Windows() []Window {
	return t.Root.windows()
}

type Node struct {
	ID          uint    `json:"id"` // The window ID, to be matched with the root focused node ID
	Window      *Window `json:"client"`
	FirstChild  *Node   `json:"firstChild"`
	SecondChild *Node   `json:"secondChild"`
}

func (n Node) windows() (windows []Window) {
	// If the node contains a window, return a tab containing information from the window
	if n.Window != nil {
		n.Window.nodeID = n.ID
		return []Window{*n.Window}
	}

	// Append tabs from child nodes
	if n.FirstChild != nil {
		windows = append(windows, n.FirstChild.windows()...)
	}
	if n.SecondChild != nil {
		windows = append(windows, n.SecondChild.windows()...)
	}
	return windows
}

// Window - A window managed by bspwm
type Window struct {
	ClassName    string `json:"className"`
	InstanceName string `json:"instanceName"`
	Active       bool   `json:"active,omitempty"`

	nodeID uint `json:"-"` // Attached ID of the node the window is a client of, used to set the active window property
}
