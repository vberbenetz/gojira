package enums

type Scope int

const (
	Scope_Authenticated Scope = iota
	Scope_Global
	Scope_Private
)
