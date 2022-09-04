package bspc

// A binary tree of nodes
type Tree struct {
	Layout        string `json:"layout"`        // Should be "monocle", otherwise ignore
	FocusedNodeID uint   `json:"focusedNodeId"` // The ID of the focused node
	Root          *Node  `json:"root"`
}

type Node struct {
	ID          uint    `json:"id"` // The window ID, to be matched with the root focused node ID
	Window      *Window `json:"client"`
	FirstChild  *Node   `json:"firstChild"`
	SecondChild *Node   `json:"secondChild"`
}

type Window struct {
	ClassName    string `json:"className"`
	InstanceName string `json:"instanceName"`
}
