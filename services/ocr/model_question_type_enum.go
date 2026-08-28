package ocr

// QuestionTypeEnum the model 'QuestionTypeEnum'
type QuestionTypeEnum string

// List of QuestionTypeEnum
const (
	QuestionTypeEnumValue0 QuestionTypeEnum = "0"
	QuestionTypeEnumValue1 QuestionTypeEnum = "1"
	QuestionTypeEnumValue2 QuestionTypeEnum = "2"
	QuestionTypeEnumValue3 QuestionTypeEnum = "3"
	QuestionTypeEnumValue4 QuestionTypeEnum = "4"
)

// All allowed values of QuestionTypeEnum enum
var AllowedQuestionTypeEnumEnumValues = []QuestionTypeEnum{
	"0",
	"1",
	"2",
	"3",
	"4",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuestionTypeEnum) IsValid() bool {
	for _, existing := range AllowedQuestionTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
