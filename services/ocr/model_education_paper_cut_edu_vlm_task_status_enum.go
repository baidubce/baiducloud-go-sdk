package ocr

// EducationPaperCutEduVlmTaskStatusEnum the model 'EducationPaperCutEduVlmTaskStatusEnum'
type EducationPaperCutEduVlmTaskStatusEnum string

// List of EducationPaperCutEduVlmTaskStatusEnum
const (
	EducationPaperCutEduVlmTaskStatusEnumPending    EducationPaperCutEduVlmTaskStatusEnum = "pending"
	EducationPaperCutEduVlmTaskStatusEnumProcessing EducationPaperCutEduVlmTaskStatusEnum = "processing"
	EducationPaperCutEduVlmTaskStatusEnumSuccess    EducationPaperCutEduVlmTaskStatusEnum = "success"
	EducationPaperCutEduVlmTaskStatusEnumFailed     EducationPaperCutEduVlmTaskStatusEnum = "failed"
)

// All allowed values of EducationPaperCutEduVlmTaskStatusEnum enum
var AllowedEducationPaperCutEduVlmTaskStatusEnumEnumValues = []EducationPaperCutEduVlmTaskStatusEnum{
	"pending",
	"processing",
	"success",
	"failed",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationPaperCutEduVlmTaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedEducationPaperCutEduVlmTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
