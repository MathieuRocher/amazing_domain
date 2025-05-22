package domain

type Course struct {
	ID          uint
	Title       string
	Description string
	Assignments []CourseAssignment
}
