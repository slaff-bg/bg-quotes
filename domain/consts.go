package domain

// Define User Ranks (still not used)
type URank uint8

const (
	URSuperAdmin URank = iota
	URAdmin
	URContributor
	URSubscriber
	URNomad
)
