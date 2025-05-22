package domain

type CourseAssignment struct {
	ID           uint
	CourseID     uint
	ClassGroupID uint
	TrainerID    uint

	Course     Course
	ClassGroup ClassGroup
	Trainer    User
}
