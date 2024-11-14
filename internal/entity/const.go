package entity

const (
	RoleUser Role = iota + 1
	RoleGroupModerator
	RoleGlobalAdmin
)

const (
	SessionKey = "session_key"
	Claims     = "claims"
)
