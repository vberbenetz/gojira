package models

import (
	"../enums"
)

type Scope struct {
	Type enums.ScopeType `json:"type"`
	Project ProjectForScope `json:"project"`
	AdditionalProperties Any `json:"additionalProperties"`
}

type DefaultShareScope struct {
	Scope enums.Scope `json:"scope"`
}
