package entity

const (
	RoleUser Role = iota + 1
	RoleGroupManager
	RoleGlobalAdmin
)

const (
	SessionKey = "session_key"
	Claims     = "claims"
)
