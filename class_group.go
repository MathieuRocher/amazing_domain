package domain

type ClassGroup struct {
	ID       uint
	Name     string
	Trainees []User
}