package domain

type Review struct {
	ID                 uint
	FormID             uint
	Form               Form
	UserID             uint
	User               User
	CourseAssignmentID uint
	CourseAssignment   CourseAssignment

	FormAnswers []FormAnswer
}
