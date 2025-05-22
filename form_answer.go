package domain

type FormAnswer struct {
	ID             uint
	ReviewID       uint
	Review         Review
	FormQuestionID uint
	FormQuestion   FormQuestion
	Option         uint
	Answer         string
	Rating         uint
}

func (f *FormAnswer) GetAnswer() interface{} {
	switch f.FormQuestion.Type {
	case Field:
		return f.Answer
	case Rating:
		return f.Rating
	case Radio:
	case Select:
		return f.Option
	default:
		return nil
	}
	return nil
}
