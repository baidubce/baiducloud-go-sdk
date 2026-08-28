package ocr

// EducationCorrectEduTaskStatusEnum the model 'EducationCorrectEduTaskStatusEnum'
type EducationCorrectEduTaskStatusEnum string

// List of EducationCorrectEduTaskStatusEnum
const (
	EducationCorrectEduTaskStatusEnumPending    EducationCorrectEduTaskStatusEnum = "pending"
	EducationCorrectEduTaskStatusEnumProcessing EducationCorrectEduTaskStatusEnum = "processing"
	EducationCorrectEduTaskStatusEnumSuccess    EducationCorrectEduTaskStatusEnum = "success"
	EducationCorrectEduTaskStatusEnumFailed     EducationCorrectEduTaskStatusEnum = "failed"
)

// All allowed values of EducationCorrectEduTaskStatusEnum enum
var AllowedEducationCorrectEduTaskStatusEnumEnumValues = []EducationCorrectEduTaskStatusEnum{
	"pending",
	"processing",
	"success",
	"failed",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationCorrectEduTaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedEducationCorrectEduTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
