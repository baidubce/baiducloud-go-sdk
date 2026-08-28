package ocr

// QuestionElementTypeEnum the model 'QuestionElementTypeEnum'
type QuestionElementTypeEnum string

// List of QuestionElementTypeEnum
const (
	QuestionElementTypeEnumValue0 QuestionElementTypeEnum = "0"
	QuestionElementTypeEnumValue1 QuestionElementTypeEnum = "1"
	QuestionElementTypeEnumValue2 QuestionElementTypeEnum = "2"
	QuestionElementTypeEnumValue3 QuestionElementTypeEnum = "3"
	QuestionElementTypeEnumValue4 QuestionElementTypeEnum = "4"
	QuestionElementTypeEnumValue5 QuestionElementTypeEnum = "5"
)

// All allowed values of QuestionElementTypeEnum enum
var AllowedQuestionElementTypeEnumEnumValues = []QuestionElementTypeEnum{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuestionElementTypeEnum) IsValid() bool {
	for _, existing := range AllowedQuestionElementTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
