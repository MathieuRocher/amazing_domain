package domain

type Form struct {
	ID            uint
	MotherId      *uint
	MotherForm    *Form
	FormQuestions []FormQuestion
}
