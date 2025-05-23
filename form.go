package domain

type Form struct {
	ID                 uint
	MotherId           *uint
	MotherForm         *Form
	CourseAssignmentId *uint
	FormQuestions      []FormQuestion
}
