package key

import "github.com/j-dumbell/lite-flag/pkg/array"

type Role string

const (
	Root     Role = "root"
	Admin    Role = "admin"
	Readonly Role = "readonly"
)

func (role Role) isValid() bool {
	return array.Includes([]Role{Root, Admin, Readonly}, role)
}
