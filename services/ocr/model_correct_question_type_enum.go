package ocr

// CorrectQuestionTypeEnum the model 'CorrectQuestionTypeEnum'
type CorrectQuestionTypeEnum int32

// List of CorrectQuestionTypeEnum
const (
	CorrectQuestionTypeEnumValue0   CorrectQuestionTypeEnum = 0
	CorrectQuestionTypeEnumValue1   CorrectQuestionTypeEnum = 1
	CorrectQuestionTypeEnumValue2   CorrectQuestionTypeEnum = 2
	CorrectQuestionTypeEnumValue3   CorrectQuestionTypeEnum = 3
	CorrectQuestionTypeEnumValue4   CorrectQuestionTypeEnum = 4
	CorrectQuestionTypeEnumValue5   CorrectQuestionTypeEnum = 5
	CorrectQuestionTypeEnumValue6   CorrectQuestionTypeEnum = 6
	CorrectQuestionTypeEnumValue7   CorrectQuestionTypeEnum = 7
	CorrectQuestionTypeEnumValue8   CorrectQuestionTypeEnum = 8
	CorrectQuestionTypeEnumValue9   CorrectQuestionTypeEnum = 9
	CorrectQuestionTypeEnumValue10  CorrectQuestionTypeEnum = 10
	CorrectQuestionTypeEnumValue11  CorrectQuestionTypeEnum = 11
	CorrectQuestionTypeEnumValue17  CorrectQuestionTypeEnum = 17
	CorrectQuestionTypeEnumValue18  CorrectQuestionTypeEnum = 18
	CorrectQuestionTypeEnumValue19  CorrectQuestionTypeEnum = 19
	CorrectQuestionTypeEnumValue401 CorrectQuestionTypeEnum = 401
	CorrectQuestionTypeEnumValue402 CorrectQuestionTypeEnum = 402
	CorrectQuestionTypeEnumValue801 CorrectQuestionTypeEnum = 801
	CorrectQuestionTypeEnumValue902 CorrectQuestionTypeEnum = 902
)

// All allowed values of CorrectQuestionTypeEnum enum
var AllowedCorrectQuestionTypeEnumEnumValues = []CorrectQuestionTypeEnum{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	17,
	18,
	19,
	401,
	402,
	801,
	902,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CorrectQuestionTypeEnum) IsValid() bool {
	for _, existing := range AllowedCorrectQuestionTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
