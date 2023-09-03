package key

import "github.com/j-dumbell/lite-flag/pkg/array"

type Role string

func (role Role) isValid() bool {
	return array.Includes([]Role{Root, Admin, Readonly}, role)
}

const (
	Root     Role = "root"
	Admin    Role = "admin"
	Readonly      = "readonly"
)
