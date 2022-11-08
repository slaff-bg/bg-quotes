package domain

// Define User Ranks (still not used)
type URank uint8

const (
	SuperAdmin URank = iota
	Admin
	Contributor
	Subscriber
	Nomad
)
