package domain

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

type User struct {
	ID           uint
	Name         string
	Email        string
	Role         Role
	Password     string
	ClassGroupID *uint
	ClassGroup   *ClassGroup
}
