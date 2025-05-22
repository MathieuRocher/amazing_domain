package domain

type FormQuestionType int

const (
	Field FormQuestionType = iota
	Rating
	Radio
	Select
)

var TypeName = map[FormQuestionType]string{
	Field:  "Field",
	Rating: "Rating",
	Radio:  "Radio",
	Select: "Select",
}

func (fqt FormQuestionType) String() string {
	return TypeName[fqt]
}

type FormQuestion struct {
	ID         uint
	FormID     uint
	Question   string
	Type       FormQuestionType
	Options    string // Array JSON
	IsRequired bool
}
