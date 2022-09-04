package bspc

// Subscribe - Subscribe to the current and future bspc state
func Subscribe() (sub chan Event) {
	sub = make(chan Event)
	return sub
}
