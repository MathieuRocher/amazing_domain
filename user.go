package domain

import "strings"

type Role int

const (
	Trainee Role = iota
	Trainer
	Administrator
)

var RoleName = map[Role]string{
	Trainee:       "Trainee",
	Trainer:       "Trainer",
	Administrator: "Admin",
}

func (r Role) String() string {
	return RoleName[r]
}

func ParseRole(input string) Role {
	for r, name := range RoleName {
		if strings.EqualFold(name, input) {
			return r
		}
	}
	return -1
}

type User struct {
	ID           uint
	Name         string
	Email        string
	Role         Role
	Password     string
	ClassGroupID *uint
	ClassGroup   *ClassGroup
}
