package key

import (
	"slices"
)

type Role string

const (
	Root     Role = "root"
	Admin    Role = "admin"
	Readonly Role = "readonly"
)

func (role Role) isValid() bool {
	return slices.Contains([]Role{Root, Admin, Readonly}, role)
}
