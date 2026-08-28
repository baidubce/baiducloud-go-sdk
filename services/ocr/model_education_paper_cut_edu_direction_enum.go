package ocr

// EducationPaperCutEduDirectionEnum the model 'EducationPaperCutEduDirectionEnum'
type EducationPaperCutEduDirectionEnum int32

// List of EducationPaperCutEduDirectionEnum
const (
	EducationPaperCutEduDirectionEnumValueMinus1 EducationPaperCutEduDirectionEnum = -1
	EducationPaperCutEduDirectionEnumValue0      EducationPaperCutEduDirectionEnum = 0
	EducationPaperCutEduDirectionEnumValue1      EducationPaperCutEduDirectionEnum = 1
	EducationPaperCutEduDirectionEnumValue2      EducationPaperCutEduDirectionEnum = 2
	EducationPaperCutEduDirectionEnumValue3      EducationPaperCutEduDirectionEnum = 3
)

// All allowed values of EducationPaperCutEduDirectionEnum enum
var AllowedEducationPaperCutEduDirectionEnumEnumValues = []EducationPaperCutEduDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationPaperCutEduDirectionEnum) IsValid() bool {
	for _, existing := range AllowedEducationPaperCutEduDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
