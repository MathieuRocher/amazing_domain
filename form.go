package domain

type Form struct {
	ID                 uint
	MotherId           *uint
	MotherForm         *Form
	CourseAssignmentId *uint
	CourseAssignment   *CourseAssignment
	FormQuestions      []FormQuestion
}
